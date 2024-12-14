package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"xevion.dev/quickvast/api"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Search() []api.ScoredOffer {
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

	return api.ScoreOffers(resp.Offers)
}
