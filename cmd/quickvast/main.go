package main

import (
	"os"

	"go.uber.org/zap"
	"xevion.dev/quickvast/api"

	"github.com/joho/godotenv"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugar := logger.Sugar()
	// Load .env file
	if err := godotenv.Load(); err != nil {
		sugar.Fatal(err)
	}

	// Get API key from environment
	apiKey := os.Getenv("VASTAI_API_KEY")
	if apiKey == "" {
		sugar.Fatal("VASTAI_API_KEY not found in environment")
	}

	// Create client
	client := api.NewClient(apiKey)

	// Create search
	search := api.NewSearch()
	// search.CPUCores = api.Ge(8)

	// Perform search
	sugar.Infow("Searching", "search", search)
	resp, err := client.Search(search)
	if err != nil {
		sugar.Fatal(err)
	}

	// Print offers
	sugar.Infof("Offers: %d", len(resp.Offers))
	for _, offer := range resp.Offers {
		sugar.Info(offer.String())
	}

	sugar.Info("Done")
}
