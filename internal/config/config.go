package config

import (
	"os"
)

type (
	Config struct {
		ServerAddress string
	}
)

const (
	serverAddressEnvKey     = "SERVER_ADDRESS"
	serverAddressEnvDefault = "0.0.0.0:8080"
)

func getEnv(envKey, envDefault string) string {
	if envValue, found := os.LookupEnv(envKey); found {
		return envValue
	}
	return envDefault
}
