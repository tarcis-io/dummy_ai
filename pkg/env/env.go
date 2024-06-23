package env

import (
	"os"
)

var (
	ServerAddress = get("SERVER_ADDRESS", ":3000")
)

func get(key string, defaultValue string) string {

	if value, isPresent := os.LookupEnv(key); isPresent {

		return value
	}

	return defaultValue
}
