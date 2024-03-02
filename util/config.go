package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config is a struct that holds all configurations for the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	ServerName           string        `mapstructure:"SERVER_NAME"`
	DBSource             string        `mapstructure:"DB_SOURCE_TOKEN_SERVICE"`
	DBSourceLocal        string        `mapstructure:"DB_SOURCE_TOKEN_SERVICE_LOCAL"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS_TOKEN_SERVICE"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS_TOKEN_SERVICE"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CertPem              string        `mapstructure:"CERT_PEM"`
	KeyPem               string        `mapstructure:"KEY_PEM"`
	CaCertPem            string        `mapstructure:"CA_CERT_PEM"`
}

// LoadConfig loads the configuration from the given path.
func LoadConfig() (config Config, err error) {
	// AutomaticEnv makes Viper check if environment variables match any of the existing keys.
	// If matching env vars are found, they are loaded into Viper.
	viper.AutomaticEnv()

	// Check if environment variables are set
	if serverName := viper.GetString("SERVER_NAME"); serverName != "" {
		config.ServerName = serverName
	}
	if dbSource := viper.GetString("DB_SOURCE_TOKEN_SERVICE"); dbSource != "" {
		config.DBSource = dbSource
	}
	if dbSourceLocal := viper.GetString("DB_SOURCE_TOKEN_SERVICE_LOCAL"); dbSourceLocal != "" {
		config.DBSourceLocal = dbSourceLocal
	}
	if migrationURL := viper.GetString("MIGRATION_URL"); migrationURL != "" {
		config.MigrationURL = migrationURL
	}
	if httpServerAddress := viper.GetString("HTTP_SERVER_ADDRESS_TOKEN_SERVICE"); httpServerAddress != "" {
		config.HttpServerAddress = httpServerAddress
	}
	if grpcServerAddress := viper.GetString("GRPC_SERVER_ADDRESS_TOKEN_SERVICE"); grpcServerAddress != "" {
		config.GrpcServerAddress = grpcServerAddress
	}
	if tokenSymmetricKey := viper.GetString("TOKEN_SYMMETRIC_KEY"); tokenSymmetricKey != "" {
		config.TokenSymmetricKey = tokenSymmetricKey
	}
	if accessTokenDuration := viper.GetDuration("ACCESS_TOKEN_DURATION"); accessTokenDuration != 0 {
		config.AccessTokenDuration = accessTokenDuration
	}
	if refreshTokenDuration := viper.GetDuration("REFRESH_TOKEN_DURATION"); refreshTokenDuration != 0 {
		config.RefreshTokenDuration = refreshTokenDuration
	}
	if certPem := viper.GetString("CERT_PEM"); certPem != "" {
		config.CertPem = certPem
	}
	if keyPem := viper.GetString("KEY_PEM"); keyPem != "" {
		config.KeyPem = keyPem
	}
	if caCertPem := viper.GetString("CA_CERT_PEM"); caCertPem != "" {
		config.CaCertPem = caCertPem
	}
	// If environment variables are not set, attempt to load from the config file
	if (config.DBSource == "" || config.DBSourceLocal == ""  || config.HttpServerAddress == "" ||
		config.GrpcServerAddress == "" || config.CertPem == "" || config.KeyPem == "" || config.CaCertPem == "") &&
		viper.GetString("CI_ENV") != "true" && viper.GetString("CONTAINER_ENV") != "true" {

		// Load configuration from the app.env file
		viper.SetConfigFile("env/app.env")
		if err := viper.ReadInConfig(); err != nil {
			return config, err
		}

		// Unmarshal the configuration into the Config struct
		if err := viper.Unmarshal(&config); err != nil {
			return config, err
		}

		return config, nil
	}

	return config, err
}
