package config

import (
	"fmt"
	"os"
)

var (
	// DB envs with default values
	DB_HOST = getEnv("DB_HOST", "db")
	DB_USER = getEnv("DB_USER", "postgres")
	DB_PASS = getEnv("DB_PASS", "postgres")
	DB_NAME = getEnv("DB_NAME", "postgres")
	DB_PORT = getEnv("DB_PORT", "5432")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
