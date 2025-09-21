package config

import (
	"os"
)

type (
	Config struct {
		ServerAddress string
	}
)

func getEnv(envKey, envDefault string) string {
	if envValue, found := os.LookupEnv(envKey); found {
		return envValue
	}
	return envDefault
}
