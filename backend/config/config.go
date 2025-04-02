package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Firebase FirebaseConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Env  string
}

// FirebaseConfig holds Firebase configuration
type FirebaseConfig struct {
	CredentialsFile string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	var cfg Config

	// Server config
	cfg.Server.Port = getEnv("PORT", "8080")
	cfg.Server.Env = getEnv("ENV", "development")

	// Firebase config
	cfg.Firebase.CredentialsFile = getEnv("FIREBASE_CREDENTIALS_FILE", "./config/firebase-credentials.json")

	return &cfg, nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
