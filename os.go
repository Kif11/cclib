package cclib

import (
	"os"
	"strconv"
)

// GetEnv get environmental variable with fallback
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvInt get environmental string and try to convert it to an integer
func GetEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		v, _ := strconv.Atoi(value)
		return v
	}
	return fallback
}
