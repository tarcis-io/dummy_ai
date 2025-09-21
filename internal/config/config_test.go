package config

import (
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name              string
		envValues         map[string]string
		wantServerAddress string
		wantError         bool
	}{}
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
			if config.ServerAddress != testCase.wantServerAddress {
				t.Errorf("New() *Config.ServerAddress got=%q want=%q", config.ServerAddress, testCase.wantServerAddress)
			}
		})
	}
}
