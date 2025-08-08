package env

import (
	"fmt"
	"net"
	"os"
	"sync"
)

type (
	Config struct {
		serverAddress string
	}
)

func (c *Config) ServerAddress() string {
	return c.serverAddress
}

var (
	appConfig *Config

	appConfigOnce sync.Once
)

func AppConfig() *Config {
	appConfigOnce.Do(func() {
		appConfig = &Config{
			serverAddress: lookupServerAddressEnv(),
		}
	})
	return appConfig
}

func lookupServerAddressEnv() string {
	v := lookupEnv("SERVER_ADDRESS", ":8080")
	_, _, err := net.SplitHostPort(v)
	if err != nil {
		panic(fmt.Errorf("[error] Failed to lookup server address: %s", err))
	}
	return v
}

func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
