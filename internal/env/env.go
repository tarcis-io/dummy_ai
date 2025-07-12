package env

import (
	"os"
)

var (
	language      = env("LANGUAGE", "en")
	serverAddress = env("SERVER_ADDRESS", ":8080")
)

func Language() string {
	return language
}

func ServerAddress() string {
	return serverAddress
}

func env(key string, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}
