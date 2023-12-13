package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"k8s.io/client-go/util/homedir"
)

var BasePathFS = getEnvWithDefault("BASE_PATH_FS", "./uploads/")

func GetJWTKey() string {
	return getEnvWithDefault("JWT_KEY", "randomkey")
}

// TODO : Follow singelton for configs
type ServerConfig struct {
	address      string
	port         string
	readTimeout  int
	writeTimeout int
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

func (c ServerConfig) GetReadTimeout() int {
	return c.readTimeout
}

func (c ServerConfig) GetWriteTimeout() int {
	return c.writeTimeout
}

func NewServerConfig() ServerConfig {
	rTimeout, errR := strconv.Atoi(getEnvWithDefault("READ_TIMEOUT", "5"))
	wTimeout, errW := strconv.Atoi(getEnvWithDefault("WRITE_TIMEOUT", "5"))

	if errR != nil || errW != nil {
		log.Fatal(errR, errW)
	}

	return ServerConfig{
		address:      getEnvWithDefault("SERVER_ADDRESS", "localhost"),
		port:         getEnvWithDefault("SERVER_PORT", "8080"),
		readTimeout:  rTimeout,
		writeTimeout: wTimeout,
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

var K8sConfig = getEnvWithDefault("KUBE_CONFIG", filepath.Join(homedir.HomeDir(), ".kube", "config"))
