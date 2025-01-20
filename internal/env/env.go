package env

var (
	serverAddress = env("SERVER_ADDRESS", ":8080")
)

func ServerAddress() string {

	return serverAddress
}

func env(key string, defaultValue string) string {

	return defaultValue
}
