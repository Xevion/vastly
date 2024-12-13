package main

import (
	"fmt"
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

	// Get instances
	resp, err := client.GetInstances()
	if err != nil {
		log.Fatalf("Error getting instances: %v", err)
	}

	if len(resp.Instances) == 0 {
		fmt.Println("No instances found")
		return
	}

	// Print instances
	for _, instance := range resp.Instances {
		fmt.Printf("Instance %d: %+v\n", instance.ID, instance)
	}
}
