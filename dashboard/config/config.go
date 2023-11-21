package config

import (
	"fmt"
	"os"
)

var BasePathFS = getEnvWithDefault("BASE_PATH_FS", "./uploads/")

// TODO : Follow singelton for configs
type ServerConfig struct {
	address string
	port    string
}

type DBConfig struct {
	user     string
	password string
	dbName   string
	address  string
	port     string
}

func (db DBConfig) GetURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.user, db.password, db.address, db.port, db.dbName)
}

func NewDBConfig() DBConfig {
	return DBConfig{
		user:     getEnvWithDefault("DB_USERNAME", "lab_rank_user"),
		password: getEnvWithDefault("DB_PASSWORD", "lab_rank_password"),
		dbName:   getEnvWithDefault("DB_NAME", "lab_rank"),
		address:  getEnvWithDefault("DB_ADDRESS", "localhost"),
		port:     getEnvWithDefault("DB_PORT", "5432"),
	}
}

func (c ServerConfig) GetAddress() string {
	return c.address
}

func (c ServerConfig) GetPort() string {
	return c.port
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		address: getEnvWithDefault("SERVER_ADDRESS", "localhost"),
		port:    getEnvWithDefault("SERVER_PORT", "8080"),
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}

type LoggerConfig struct {
	LogLevel  string
	LogFormat string
	// Add other logger configuration options as needed
}

func InitLoggerConfig() LoggerConfig {
	return LoggerConfig{
		LogLevel:  getEnvWithDefault("LOG_LEVEL", "INFO"),  // Default log level
		LogFormat: getEnvWithDefault("LOG_FORMAT", "json"), // Default log format
		// Initialize other configuration options with default values
	}
}
