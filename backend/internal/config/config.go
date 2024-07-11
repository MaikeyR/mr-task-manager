package config

import (
	"fmt"
	"os"
)

// Config represents the database configuration settings.
type Config struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
}

// LoadConfig loads the database configuration settings from environment variables.
func LoadConfig() *Config {
		return &Config{
				DBHost:     os.Getenv("PGHOST"),
				DBPort:     os.Getenv("PGPORT"),
				DBUser:     os.Getenv("PGUSER"),
				DBPassword: os.Getenv("PGPASSWORD"),
				DBName:     os.Getenv("PGNAME"),
		}
}

// GetDBConnectionString returns the database connection string based on the loaded configuration.
func (config *Config) GetDBConnectionString() string {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
}