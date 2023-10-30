package config

import (
	"os"
)

type ServerConfig interface {
	GetAddress() string
	GetPort() string
}

type envServerConfig struct{}

func (c envServerConfig) GetAddress() string {
	if addr := os.Getenv("SERVER_ADDRESS"); addr != "" {
		return addr
	}
	return "localhost"
}

func (c envServerConfig) GetPort() string {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		return port
	}
	return "8080"
}

func NewEnvServerConfig() ServerConfig {
	return envServerConfig{}
}

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

type LoggerConfig struct {
	LogLevel  string
	LogFormat string
	// Add other logger configuration options as needed
}

func NewLoggerConfig() LoggerConfig {
	return LoggerConfig{
		LogLevel:  GetEnvWithDefault("LOG_LEVEL", "INFO"),  // Default log level
		LogFormat: GetEnvWithDefault("LOG_FORMAT", "json"), // Default log format
		// Initialize other configuration options with default values
	}
}
