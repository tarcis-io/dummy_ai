// Package config provides configurations for the application.
package config

import (
	"fmt"
	"net"
	"os"
)

type (
	// Config holds the configurations for the application.
	Config struct {
		// ServerAddress is the host and port for the server to listen on.
		ServerAddress string

		// ServerShutdownTimeout is the timeout for the server to shutdown gracefully.
		ServerShutdownTimeout int
	}
)

const (
	// Server configurations.
	serverAddressEnvKey             = "SERVER_ADDRESS"
	serverAddressEnvDefault         = "0.0.0.0:8080"
	serverShutdownTimeoutEnvKey     = "SERVER_SHUTDOWN_TIMEOUT"
	serverShutdownTimeoutEnvDefault = "10"
)

// New creates and returns a new Config instance by resolving the configurations for the application.
// It returns an error if any of the configurations cannot be resolved.
func New() (*Config, error) {
	serverAddress, err := resolveServerAddress()
	if err != nil {
		return nil, err
	}
	config := &Config{
		ServerAddress: serverAddress,
	}
	return config, nil
}

// resolveServerAddress resolves the server address configuration for the application.
// It returns an error if the configuration cannot be resolved.
func resolveServerAddress() (string, error) {
	serverAddress := getEnv(serverAddressEnvKey, serverAddressEnvDefault)
	if _, _, err := net.SplitHostPort(serverAddress); err != nil {
		return "", fmt.Errorf("invalid server address %s=%q error=%w", serverAddressEnvKey, serverAddress, err)
	}
	return serverAddress, nil
}

// getEnv retrieves the value of the environment variable with the given key.
// It returns the default value if the environment variable is not set.
func getEnv(envKey, envDefault string) string {
	if envValue, found := os.LookupEnv(envKey); found {
		return envValue
	}
	return envDefault
}
