package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://testuser:testpassword@localhost/computerdb?sslmode=disable"),
		Port:        getEnv("PORT", "3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
