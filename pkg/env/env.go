package env

import (
	"os"
)

var (
	serverAddress = get("SERVER_ADDRESS", ":3000")
)

func ServerAddress() string {

	return serverAddress
}

func get(key string, defaultValue string) string {

	if value, isPresent := os.LookupEnv(key); isPresent {

		return value
	}

	return defaultValue
}
