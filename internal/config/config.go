package config

import (
	"os"
)

func getEnv(envKey, envDefault string) string {
	if envValue, found := os.LookupEnv(envKey); found {
		return envValue
	}
	return envDefault
}
