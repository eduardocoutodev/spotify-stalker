package config

import (
	"os"
)

// TODO Convert this to generic type
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
