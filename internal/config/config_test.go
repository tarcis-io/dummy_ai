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
