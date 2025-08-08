// Package env provides access to environment variables.
package env

import (
	"os"
	"sync"
)

type (
	// Config represents the configuration settings for the application.
	Config struct {
		languageCode  string
		serverAddress string
		themeCode     string
	}
)

// LanguageCode returns the language code for the application.
func (c *Config) LanguageCode() string {
	return c.languageCode
}

// ServerAddress returns the server address for the application.
func (c *Config) ServerAddress() string {
	return c.serverAddress
}

// ThemeCode returns the theme code for the application.
func (c *Config) ThemeCode() string {
	return c.themeCode
}

var (
	// currentConfig holds the current configuration settings for the application.
	currentConfig *Config

	// currentConfigOnce is used to ensure that the current configuration settings
	// are only initialized once.
	currentConfigOnce sync.Once
)

// CurrentConfig returns the current configuration settings for the application.
func CurrentConfig() *Config {
	currentConfigOnce.Do(func() {
		currentConfig = &Config{
			languageCode:  lookupEnv("LANGUAGE_CODE", "en"),
			serverAddress: lookupEnv("SERVER_ADDRESS", ":8080"),
			themeCode:     lookupEnv("THEME_CODE", "dark"),
		}
	})
	return currentConfig
}

// lookupEnv retrieves the value of the environment variable named by the key
// or the default value if the variable is not present.
func lookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
