package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"xevion.dev/vastly/api"
)

// App struct
type App struct {
	ctx     context.Context
	client  *api.Client
	logger  *zap.SugaredLogger
	latency *api.LatencyQueue
}

// NewApp creates a new App application struct
func NewApp() *App {
	logger, _ := zap.NewDevelopment()
	return &App{
		logger:  logger.Sugar(),
		latency: api.NewLatencyQueue(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load .env file
	if err := godotenv.Load(); err != nil {
		a.logger.Fatal(err)
	}

	// Get API key from environment
	apiKey := os.Getenv("VASTAI_API_KEY")
	if apiKey == "" {
		a.logger.Fatal("VASTAI_API_KEY not found in environment")
	}
	a.client = api.NewClient(apiKey)

	// Start latency queue
	go a.latency.Start()
}

func (a *App) beforeClose(ctx context.Context) bool {
	a.latency.Kill()
	return false
}

func (a *App) Search() []api.ScoredOffer {
	defer a.logger.Sync()

	// Create search
	search := api.NewSearch()
	search.AllocatedStorage = 39.94657756485159
	search.Limit = 1000

	// Perform search
	a.logger.Infow("Searching", "search", search)
	resp, err := a.client.Search(search)
	if err != nil {
		a.logger.Fatal(err)
	}

	return api.ScoreOffers(resp.Offers)
}
