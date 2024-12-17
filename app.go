package main

import (
	"context"
	"fmt"
	"os"
	"time"

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
	redis   *redis.Client
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

	a.redis = redis.NewClient(redisOptions)
	if _, err := a.redis.Ping(ctx).Result(); err != nil {
		a.logger.Fatal("Failed to connect to Redis", err)
	}

	// Start latency queue
	a.latency = api.NewLatencyQueue(a.redis)
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

	scored := api.ScoreOffers(resp.Offers)

	// collect IP latency keys
	currentIP := a.latency.GetSelfIP()
	batchCount := min(100, len(scored))
	latencyKeys := make([]string, batchCount, batchCount)
	for i := range batchCount {
		latencyKeys[i] = fmt.Sprintf("latency:%s:%s", currentIP, scored[i].Offer.PublicIPAddr)
	}
	latencyValues, err := a.redis.MGet(a.ctx, latencyKeys...).Result()
	if err != nil {
		a.logger.Errorw("Failed to get latency values", "error", err)
	}

	// Assign latency values to scored offers
	for i, value := range latencyValues {
		if value != nil {
			if value.(string) == "timeout" {
				scored[i].Latency = api.Pointer(int32(-1))
				continue
			}

			parsed, err := time.ParseDuration(value.(string))
			if err != nil {
				a.logger.Errorw("Failed to parse latency value", "error", err)
				continue
			}
			scored[i].Latency = api.Pointer(int32(parsed.Milliseconds()))
		}
	}

	return scored
}
