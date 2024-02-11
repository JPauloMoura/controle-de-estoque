package configs

import (
	"log"
	"os"
)

// BuildConfig imports the necessary environment variables and makes them available in a config structure
func BuildConfig() *Config {
	cfg := Config{
		dbConnectionStr: os.Getenv("DB_CONNECTION_STRING"),
		serverPort:      os.Getenv("SERVER_PORT"),
		logType:         os.Getenv("LOG_TYPE"),
		jwtKey:          os.Getenv("JWT_KEY"),
	}

	cfg.validate()

	return &cfg
}

// Config contains the application variables
type Config struct {
	dbConnectionStr string
	serverPort      string
	logType         string
	jwtKey          string
}

func (c Config) DbConnectionStr() string { return c.dbConnectionStr }
func (c Config) ServerPort() string      { return c.serverPort }
func (c Config) LogType() string         { return c.logType }
func (c Config) JwtKey() string          { return c.jwtKey }

func (c Config) validate() {
	if c.dbConnectionStr == "" {
		log.Fatal("DB_CONNECTION_STRING is required")
	}
	if c.serverPort == "" {
		log.Fatal("SERVER_PORT is required")
	}
	if c.jwtKey == "" {
		log.Fatal("JWT_KEY is required")
	}
}
