package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds API required configs
type Config struct {
	MySQL *MySQLConfig
}

// MySQLConfig config for mysql
type MySQLConfig struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

// ApplicationConfig application config
var ApplicationConfig *Config

// Setup function to setup config
func Setup() {
	ApplicationConfig = new(Config)
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed loading .env")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlHost != "" && mysqlDatabase != "" && mysqlUser != "" && mysqlPassword != "" && mysqlPort != "" {
		mysqlConfig := MySQLConfig{
			Host:     mysqlHost,
			Database: mysqlDatabase,
			User:     mysqlUser,
			Password: mysqlPassword,
			Port:     mysqlPort,
		}
		ApplicationConfig.MySQL = &mysqlConfig
	}
}
