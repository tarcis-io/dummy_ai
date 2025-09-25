// Package main is the entry point for the dummy_ai application.
package main

import (
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	return nil
}
