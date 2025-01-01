package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	valIsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valIsInt
}
