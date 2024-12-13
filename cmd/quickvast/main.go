package main

import (
	"log"
	"os"

	"xevion.dev/quickvast/api"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get API key from environment
	apiKey := os.Getenv("VASTAI_API_KEY")
	if apiKey == "" {
		log.Fatal("VASTAI_API_KEY not found in environment")
	}

	// Create client
	client := api.NewClient(apiKey)
}
