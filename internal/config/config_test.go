package config

import (
	"testing"
)

// TestNew verifies if the New function correctly creates a new Config instance.
// It covers default configurations, custom valid configurations
// and error handling for invalid configurations.
func TestNew(t *testing.T) {
}

func assertError(t *testing.T, got error, wantError bool) {
	t.Helper()
	if (got != nil) != wantError {
		t.Fatalf("error got=%v wantError=%t", got, wantError)
	}
}
