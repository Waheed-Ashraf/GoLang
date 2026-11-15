package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBPort     string
	DBName     string
	DBSSLMode  string
	DBHost     string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port: getEnv("PORT" , "8080"),
		DBUser: getEnv("DB_USER" , "postgres"),
		DBHost: getEnv("DB_HOST" , "localhost"),
		DBPort: getEnv("DB_PORT" , "5432"),
		DBName: getEnv("DB_NAME" , "ginapidemo"),
		DBPassword: getEnv("DB_PASSWORD" , "12345678"),
		DBSSLMode: getEnv("DBSSLMODE" , "disable"),
	}

	if cfg.DBUser == "" || cfg.DBName == "" { 
		return nil , errors.New("Database Name Must be valid")
	}

	return cfg , nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		c.DBHost , c.DBPort , c.DBUser , c.DBPassword , c.DBName , c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val , ok := os.LookupEnv("PORT"); ok && val != "" { 
		return val
	}
	return fallback
}
