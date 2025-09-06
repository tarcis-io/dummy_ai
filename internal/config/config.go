package config

type (
	Config struct {
		ServerAddress string
	}
)

const (
	serverAddressKey     = "SERVER_ADDRESS"
	serverAddressDefault = "0.0.0.0:8080"
)
