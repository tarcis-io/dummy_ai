package config

import (
	"reflect"
	"testing"
	"time"
)

// TestNew verifies if the New function correctly creates a new Config instance.
// It covers default configurations, custom valid configurations
// and error handling for invalid configurations.
func TestNew(t *testing.T) {
	mustParseDuration := func(duration string) time.Duration {
		timeDuration, err := time.ParseDuration(duration)
		if err != nil {
			panic(err)
		}
		return timeDuration
	}
	testCases := []struct {
		name       string
		envValues  map[string]string
		wantConfig *Config
		wantError  bool
	}{
		{
			name: "should create a new Config instance with default values",
			wantConfig: &Config{
				ServerAddress:         serverAddressEnvDefault,
				ServerShutdownTimeout: mustParseDuration(serverShutdownTimeoutEnvDefault),
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom server address: 127.0.0.1:8081",
			envValues: map[string]string{
				serverAddressEnvKey: "127.0.0.1:8081",
			},
			wantConfig: &Config{
				ServerAddress:         "127.0.0.1:8081",
				ServerShutdownTimeout: mustParseDuration(serverShutdownTimeoutEnvDefault),
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom server address: localhost:3000",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost:3000",
			},
			wantConfig: &Config{
				ServerAddress:         "localhost:3000",
				ServerShutdownTimeout: mustParseDuration(serverShutdownTimeoutEnvDefault),
			},
			wantError: false,
		},
		{
			name: "should return an error if the server address cannot be resolved: empty",
			envValues: map[string]string{
				serverAddressEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server address cannot be resolved: 127.0.0.1",
			envValues: map[string]string{
				serverAddressEnvKey: "127.0.0.1",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server address cannot be resolved: localhost:3000:",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost:3000:",
			},
			wantError: true,
		},
		{
			name: "should create a new Config instance with custom server shutdown timeout: 15s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "15s",
			},
			wantConfig: &Config{
				ServerAddress:         serverAddressEnvDefault,
				ServerShutdownTimeout: mustParseDuration("15s"),
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
			name: "should return an error if the server shutdown timeout cannot be resolved: empty",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be resolved: invalid",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "invalid",
			},
			wantError: true,
		},
		{
			name: "should return an error if the server shutdown timeout cannot be resolved: -1s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "-1s",
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
		{
			name: "should create a new Config instance with custom values",
			envValues: map[string]string{
				serverAddressEnvKey:         ":9090",
				serverShutdownTimeoutEnvKey: "10m",
			},
			wantConfig: &Config{
				ServerAddress:         ":9090",
				ServerShutdownTimeout: mustParseDuration("10m"),
			},
			wantError: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for envKey, envValue := range testCase.envValues {
				t.Setenv(envKey, envValue)
			}
			config, err := New()
			if (err != nil) != testCase.wantError {
				t.Fatalf("New() error got=%v wantError=%v", err, testCase.wantError)
			}
			if !reflect.DeepEqual(config, testCase.wantConfig) {
				t.Fatalf("New() *Config\ngot=%#v\nwant=%#v", config, testCase.wantConfig)
			}
		})
	}
}
