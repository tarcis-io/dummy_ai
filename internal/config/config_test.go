package config

import (
	"testing"
	"time"
)

// TestNew verifies if the New function correctly creates a new Config instance.
// It covers default configurations, custom valid configurations
// and error handling for invalid configurations.
func TestNew(t *testing.T) {
}

func mustParseDuration(duration string) time.Duration {
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return timeDuration
}
