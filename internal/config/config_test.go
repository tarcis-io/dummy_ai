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
		})
	}
}
