package env

import (
	"os"
)

type (
	Config struct {
		languageCode  string
		serverAddress string
		themeCode     string
	}
)

func (c *Config) LanguageCode() string {
	return c.languageCode
}

func (c *Config) ServerAddress() string {
	return c.serverAddress
}

func (c *Config) ThemeCode() string {
	return c.themeCode
}

var (
	currentConfig *Config
)

func init() {
	currentConfig = &Config{
		languageCode:  LookupEnv("LANGUAGE_CODE", "en"),
		serverAddress: LookupEnv("SERVER_ADDRESS", ":8080"),
		themeCode:     LookupEnv("THEME_CODE", "dark"),
	}
}

func CurrentConfig() *Config {
	return currentConfig
}

func LanguageCode() string {
	return currentConfig.languageCode
}

func ServerAddress() string {
	return currentConfig.serverAddress
}

func ThemeCode() string {
	return currentConfig.themeCode
}

func LookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
