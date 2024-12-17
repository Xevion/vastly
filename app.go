package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
		logger: logger.Sugar(),
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

	// Create Vast API client
	a.client = api.NewClient(apiKey)

	// Connect to Redis
	redisUrl := os.Getenv("REDIS_URL")
	redisOptions, err := redis.ParseURL(redisUrl)
	if err != nil {
		a.logger.Fatal("Failed to parse Redis URL", err)
	}
	redis := redis.NewClient(redisOptions)

	// Start latency queue
	a.latency = api.NewLatencyQueue(redis)
	go a.latency.Start(ctx)
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

	for _, offer := range resp.Offers {
		a.latency.QueuePing(offer.PublicIPAddr)
	}

	return api.ScoreOffers(resp.Offers)
}
