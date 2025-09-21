package config

import (
	"fmt"
	"net"
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

func resolveServerAddress() (string, error) {
	serverAddress := getEnv(serverAddressEnvKey, serverAddressEnvDefault)
	if _, _, err := net.SplitHostPort(serverAddress); err != nil {
		return "", fmt.Errorf("invalid server address %s=%q error=%w", serverAddressEnvKey, serverAddress, err)
	}
	return serverAddress, nil
}

func getEnv(envKey, envDefault string) string {
	if envValue, found := os.LookupEnv(envKey); found {
		return envValue
	}
	return envDefault
}
