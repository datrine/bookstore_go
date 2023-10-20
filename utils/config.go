package utils

import (
	"os"
)

type Config struct {
	DATABASE_URL string
}

func Get() *Config {
	return &Config{
		DATABASE_URL: getEnv("DATABASE_URL"),
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}
