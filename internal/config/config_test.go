package config

import (
	"testing"
)

// TestNew verifies if the New function correctly creates a new Config instance.
func TestNew(t *testing.T) {
	testCases := []struct {
		name       string
		envValues  map[string]string
		wantConfig *Config
		wantError  bool
	}{
		{
			name:      "should create a new Config instance with default values",
			envValues: map[string]string{},
			wantConfig: &Config{
				ServerAddress: serverAddressEnvDefault,
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom server address: 127.0.0.1:8081",
			envValues: map[string]string{
				serverAddressEnvKey: "127.0.0.1:8081",
			},
			wantConfig: &Config{
				ServerAddress: "127.0.0.1:8081",
			},
			wantError: false,
		},
		{
			name: "should create a new Config instance with custom server address: localhost:3000",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost:3000",
			},
			wantConfig: &Config{
				ServerAddress: "localhost:3000",
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
			name: "should return an error if the server address cannot be resolved: localhost:999999",
			envValues: map[string]string{
				serverAddressEnvKey: "localhost:999999",
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
		})
	}
}
