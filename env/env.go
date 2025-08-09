package env

import (
	"log"
	"net"
	"os"
	"sync"
)

type (
	// Config represents the application configuration.
	Config struct {
		ServerAddress string
	}
)

const (
	serverAddressKey          = "SERVER_ADDRESS"
	serverAddressDefaultValue = ":8080"
)

var (
	// appConfig holds the application configuration.
	appConfig *Config

	// appConfigOnce is used to ensure that the application configuration
	// is initialized only once.
	appConfigOnce sync.Once
)

// AppConfig returns the application configuration.
func AppConfig() *Config {
	appConfigOnce.Do(func() {
		appConfig = &Config{
			ServerAddress: lookupServerAddressEnv(),
		}
	})
	return appConfig
}

// lookupServerAddressEnv looks up the server address environment variable.
func lookupServerAddressEnv() string {
	serverAddress := lookupEnv(serverAddressKey, serverAddressDefaultValue)
	_, _, err := net.SplitHostPort(serverAddress)
	if err != nil {
		log.Printf("[warn] Invalid value for %s %q: %v. Using default value %q.", serverAddressKey, serverAddress, err, serverAddressDefaultValue)
		return serverAddressDefaultValue
	}
	return serverAddress
}

// lookupEnv looks up an environment variable.
func lookupEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue
}
