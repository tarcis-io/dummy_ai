package env

import (
	"os"
)

var (
	serverAddress = env("SERVER_ADDRESS", ":8080")
)

func ServerAddress() string {

	return serverAddress
}

func env(key string, defaultValue string) string {

	if value, found := os.LookupEnv(key); found {

		return value
	}

	return defaultValue
}
