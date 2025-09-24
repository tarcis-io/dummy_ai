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
			name: "should create a new Config instance with custom server shutdown timeout: 20s",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "20s",
			},
			wantConfig: &Config{
				ServerAddress:         serverAddressEnvDefault,
				ServerShutdownTimeout: mustParseDuration("20s"),
			},
		},
		{
			name: "should create a new Config instance with custom server shutdown timeout: 10m",
			envValues: map[string]string{
				serverShutdownTimeoutEnvKey: "10m",
			},
			wantConfig: &Config{
				ServerAddress:         serverAddressEnvDefault,
				ServerShutdownTimeout: mustParseDuration("10m"),
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
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for envKey, envValue := range testCase.envValues {
				t.Setenv(envKey, envValue)
			}
			config, err := New()
			if testCase.wantError {
				if err == nil {
					t.Fatal("New() error got=nil want!=nil")
				}
				if config != nil {
					t.Fatal("New() *Config got!=nil want=nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("New() error got=%v want=nil", err)
			}
			if config == nil {
				t.Fatal("New() *Config got=nil want!=nil")
			}
			if !reflect.DeepEqual(config, testCase.wantConfig) {
				t.Errorf("New() *Config got=%v want=%v", config, testCase.wantConfig)
			}
		})
	}
}
