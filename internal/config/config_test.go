package config

import (
	"reflect"
	"testing"
	"time"
)

type (
	// testCase holds the test data for a single test case.
	testCase struct {
		name       string
		envValues  map[string]string
		wantConfig *Config
		wantError  bool
	}
)

// TestNew verifies if the New function correctly creates a new Config instance.
// It covers default configurations, custom valid configurations
// and error handling for invalid configurations.
func TestNew(t *testing.T) {
	testCases := []testCase{
		{
			name: "should create a new Config instance with default values",
			wantConfig: &Config{
				ServerAddress:         serverAddressEnvDefault,
				ServerShutdownTimeout: mustParseDuration(serverShutdownTimeoutEnvDefault),
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom values",
			envValues: map[string]string{
				serverAddressEnvKey:         "localhost:8081",
				serverShutdownTimeoutEnvKey: "15s",
			},
			wantConfig: &Config{
				ServerAddress:         "localhost:8081",
				ServerShutdownTimeout: mustParseDuration("15s"),
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom server address: 127.0.0.1:3000",
			envValues: map[string]string{
				serverAddressEnvKey: "127.0.0.1:3000",
			},
			wantConfig: &Config{
				ServerAddress:         "127.0.0.1:3000",
				ServerShutdownTimeout: mustParseDuration(serverShutdownTimeoutEnvDefault),
			},
		},
		{
			name: "should create a new Config instance with custom server shutdown timeout: 1m",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "1m",
			},
			wantConfig: &Config{
				ServerAddress:         serverAddressEnvDefault,
				ServerShutdownTimeout: mustParseDuration("1m"),
			},
		},
		{
			name: "should return an error if the server address cannot be resolved: empty",
			envValues: map[string]string{
				serverAddressEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server address cannot be resolved: localhost",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server address cannot be resolved: 127.0.0.1:8080:",
			envValues: map[string]string{
				serverAddressEnvKey: "127.0.0.1:8080:",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be resolved: empty",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be resolved: 15sec",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "15sec",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be resolved: -15s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "-15s",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be resolved: 0s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "0s",
			},
			wantError: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for envKey, envValue := range testCase.envValues {
				t.Setenv(envKey, envValue)
			}
			config, err := New()
			if (err != nil) != testCase.wantError {
				t.Fatalf("New() error got=%v wantError=%t", err, testCase.wantError)
			}
			if testCase.wantError {
				return
			}
			if !reflect.DeepEqual(config, testCase.wantConfig) {
				t.Fatalf("New() *Config\ngot=%#v\nwant=%#v", config, testCase.wantConfig)
			}
		})
	}
}

// mustParseDuration is a helper function that parses a duration string into a time.Duration.
// It panics if the parsing fails.
func mustParseDuration(duration string) time.Duration {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeDuration
}
