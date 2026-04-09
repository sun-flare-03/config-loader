package main

import (
	"fmt"
	"log"
	"os"
)

// config-loader — Universal config loader supporting YAML, TOML, JSON, and env vars
func main() {
	logger := log.New(os.Stdout, "[config-loader] ", log.LstdFlags)
	logger.Println("Starting application...")

	if err := run(); err != nil {
		logger.Fatalf("Fatal error: %v", err)
	}
}

func run() error {
	fmt.Println("Application initialized successfully")
	return nil
}
