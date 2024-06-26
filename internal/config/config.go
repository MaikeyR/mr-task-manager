package config

import (
	"fmt"
	"os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() *Config {
    return &Config{
        DBHost:     os.Getenv("PGHOST"),
        DBPort:     os.Getenv("PGPORT"),
        DBUser:     os.Getenv("PGUSER"),
        DBPassword: os.Getenv("PGPASSWORD"),
        DBName:     os.Getenv("PGNAME"),
    }
}

func (c *Config) GetDBConnectionString() string {
    return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}
