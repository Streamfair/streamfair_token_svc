package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	// vault "github.com/hashicorp/vault/api"
	db "github.com/Streamfair/streamfair_token_svc/db/sqlc"
	"github.com/Streamfair/streamfair_token_svc/gapi"
	"github.com/Streamfair/streamfair_token_svc/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("Hello, Streamfair Token Management Service!")
	configPath, err := filepath.Abs("app.env")
	if err != nil {
		log.Printf("config: error while getting absolute path: %v\n", err)
	}
	tlsPath, err := filepath.Abs("ssl")
	if err != nil {
		log.Printf("config: error while getting absolute path: %v\n", err)
	}

	config, err := util.LoadConfig(configPath, tlsPath)
	if err != nil {
		log.Printf("config: error while loading config: %v\n", err)
	}

	poolConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Printf("config: error while parsing config: %v\n", err)
	}

	conn, err := pgxpool.New(context.Background(), poolConfig.ConnString())
	if err != nil {
		log.Printf("db connection: unable to create connection pool: %v\n", err)
	}

	store := db.NewStore(conn)
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Printf("server: error while creating server: %v\n", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	go server.RunGrpcGatewayServer()
	server.RunGrpcServer()
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("db migration: unable to create migration: %v\n", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("db migration: unable to apply migration: %v\n", err)
	}

	log.Println("db migrated successfully")
}

// func vaulTests() {
// 	// Vault Test
// 	vaultConfig := vault.DefaultConfig()

// 	vaultConfig.Address = "http://vault_server:8200"

// 	client, err := vault.NewClient(vaultConfig)
// 	if err != nil {
// 		log.Fatalf("unable to initialize Vault client: %v", err)
// 	}

// 	client.SetToken("dev-only-token")

// 	// Write a secret
// 	secretData := map[string]interface{}{
// 		"password": "Hashi123",
// 	}

// 	_, err = client.KVv2("secret").Put(context.Background(), "my-secret-password", secretData)
// 	if err != nil {
// 		log.Fatalf("unable to write secret: %v", err)
// 	}

// 	fmt.Println("Secret written successfully.")

// 	// Read the secret
// 	secret, err := client.KVv2("secret").Get(context.Background(), "my-secret-password")
// 	if err != nil {
// 		log.Fatalf("unable to read secret: %v", err)
// 	}

// 	value, ok := secret.Data["password"].(string)
// 	if !ok {
// 		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
// 	}

// 	// Check the value
// 	if value != "Hashi123" {
// 		log.Fatalf("unexpected password value %q retrieved from vault", value)
// 	}

// 	fmt.Println("Access granted!")

// 	// End of Vault Test
// }
