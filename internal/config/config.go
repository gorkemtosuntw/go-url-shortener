package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig holds all server related configuration
type ServerConfig struct {
	Port     string
	BaseURL  string
	Protocol string
	Domain   string
}

// DatabaseConfig holds all database related configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:     getEnv("SERVER_PORT", "8080"),
			Domain:   getEnv("DOMAIN", "localhost"),
			Protocol: getEnv("PROTOCOL", "http"),
			BaseURL:  getEnv("BASE_URL", "http://localhost:8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     5432, // Default PostgreSQL port
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "1"),
			DBName:   getEnv("DB_NAME", "url_shortener"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
