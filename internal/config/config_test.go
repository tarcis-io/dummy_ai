package config

import (
	"testing"
	"time"
)

// TestNew verifies if the New function correctly creates a new Config instance.
func TestNew(t *testing.T) {
	serverShutdownTimeoutDefault, _ := time.ParseDuration(serverShutdownTimeoutEnvDefault)
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
				ServerShutdownTimeout: serverShutdownTimeoutDefault,
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
				ServerShutdownTimeout: serverShutdownTimeoutDefault,
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
				ServerShutdownTimeout: serverShutdownTimeoutDefault,
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
			name: "should return an error if the server address cannot be resolved: localhost",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost",
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
			if config.ServerAddress != testCase.wantConfig.ServerAddress {
				t.Errorf("New() *Config.ServerAddress got=%q want=%q", config.ServerAddress, testCase.wantConfig.ServerAddress)
			}
			if config.ServerShutdownTimeout != testCase.wantConfig.ServerShutdownTimeout {
				t.Errorf("New() *Config.ServerShutdownTimeout got=%q want=%q", config.ServerShutdownTimeout, testCase.wantConfig.ServerShutdownTimeout)
			}
		})
	}
}
