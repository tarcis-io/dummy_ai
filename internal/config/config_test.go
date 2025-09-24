package config

import (
	"testing"
	"time"
)

type (
	// testCase
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
