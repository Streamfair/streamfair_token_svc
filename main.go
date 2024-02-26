package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

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
