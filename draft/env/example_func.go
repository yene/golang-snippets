//go:build exclude
// +build exclude

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// first: load .env into environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	deployEnvironment, err := parseEnv("DEPLOY_PORT", "8081", false)
	if err != nil {
		log.Fatalf("Error in configuration: %v", err) // How to print error string nicely here
	}
}

// parseEnv is very opinionated about required and empty behaviour
func parseEnv(name string, defaultValue string, isRequired bool) (string, error) {
	val, present := os.LookupEnv(name)
	if !present && isRequired {
		return "", fmt.Errorf("required environment variable %s is not set", name)
	}
	if !present || val == "" {
		val = defaultValue
	}
	return val, nil
}
