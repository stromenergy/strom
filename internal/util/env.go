package util

import (
	"os"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	
	if !exists {
		value = fallback
	}

	return value
}

func GetEnvBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	return ParseBool(value, fallback)
}

func GetEnvFloat64(key string, fallback float64) float64 {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	return ParseFloat64(value, fallback)
}

func GetEnvInt32(key string, fallback int32) int32 {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	return ParseInt32(value, fallback)
}