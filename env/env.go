package env

import (
	"fmt"
	"net"
	"os"
	"sync"
)

type (
	// Config represents the application configuration.
	Config struct {
		serverAddress string
	}
)

// ServerAddress returns the server address configuration.
func (c *Config) ServerAddress() string {
	return c.serverAddress
}

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
			serverAddress: lookupServerAddressEnv(),
		}
	})
	return appConfig
}

// lookupServerAddressEnv looks up the server address environment variable.
func lookupServerAddressEnv() string {
	v := lookupEnv("SERVER_ADDRESS", ":8080")
	_, _, err := net.SplitHostPort(v)
	if err != nil {
		panic(fmt.Errorf("[error] Failed to lookup server address: %s", err))
	}
	return v
}

// lookupEnv looks up an environment variable.
func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
