// Package env provides access to environment variables.
package env

import (
	"os"
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
)

// init runs on startup.
// It configures the package's initial state.
// 1. The current configuration.
func init() {
	currentConfig = &Config{
		languageCode:  LookupEnv("LANGUAGE_CODE", "en"),
		serverAddress: LookupEnv("SERVER_ADDRESS", ":8080"),
		themeCode:     LookupEnv("THEME_CODE", "dark"),
	}
}

// CurrentConfig returns the current configuration settings for the application.
func CurrentConfig() *Config {
	return currentConfig
}

// LanguageCode returns the language code for the application.
func LanguageCode() string {
	return currentConfig.languageCode
}

// ServerAddress returns the server address for the application.
func ServerAddress() string {
	return currentConfig.serverAddress
}

// ThemeCode returns the theme code for the application.
func ThemeCode() string {
	return currentConfig.themeCode
}

// LookupEnv retrieves the value of the environment variable named by the key
// or the default value if the variable is not present.
func LookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
