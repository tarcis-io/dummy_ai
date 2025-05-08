package env

import (
	"os"
)

func env(key string, defaultValue string) string {

	if value, found := os.LookupEnv(key); found {

		return value
	}

	return defaultValue
}
