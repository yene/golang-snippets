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

	deployEnvironment := os.Getenv("DEPLOY_ENVIRONMENT")
	deployPort := os.Getenv("DEPLOY_PORT")

	fmt.Println("Environment:", deployEnvironment)
	fmt.Println("Port:", deployPort)
}
