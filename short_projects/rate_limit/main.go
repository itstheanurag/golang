package main

import (
	"context"
	"fmt"
	"net/http"
	"rate-limit/middleware"
	"rate-limit/ratelimit"
	"time"

	"github.com/redis/go-redis/v9"
)

func Handler(limiterName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "Hello from %s endpoint!", "time": "%s"}`+"\n", limiterName, time.Now().Format(time.RFC3339))
	}
}

func main() {
	fmt.Println("RATE LIMITING TEST SERVER")

	// 1. Fixed Window Limiter: 10 requests per 20 seconds
	fixedLimiter := ratelimit.NewFixedWindowRateLimiter(10, 20*time.Second)

	// 2. Sliding Window Limiter (Redis): 10 requests per 20 seconds
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	slidingLimiter := ratelimit.NewSlidingWindowRateLimiter(rdb, 10, 20*time.Second)

	// 3. Token Bucket Limiter: 5 tokens per second, capacity 10
	tokenLimiter := ratelimit.NewTokenBucketRateLimiter(5, 10)

	// 4. Client Tier Limiter
	tierLimiter := ratelimit.NewClientTierRateLimiter(ratelimit.TierConfig{Rate: 1, Capacity: 5}) // Default: 1 req/sec
	tierLimiter.AddTier("gold", ratelimit.TierConfig{Rate: 10, Capacity: 20})                     // Gold: 10 req/sec
	tierLimiter.SetClientTier("api_key:gold-member", "gold")

	// Register Routes
	http.Handle("/fixed", middleware.RateLimiterMiddleware(fixedLimiter, Handler("Fixed Window")))
	
	// Sliding window implementation uses context, but the middleware interface only takes string.
	// We need to wrap it to match the interface.
	http.Handle("/sliding", middleware.RateLimiterMiddleware(wrapSlidingLimiter(slidingLimiter), Handler("Sliding Window")))
	
	http.Handle("/token", middleware.RateLimiterMiddleware(tokenLimiter, Handler("Token Bucket")))
	http.Handle("/tier", middleware.RateLimiterMiddleware(tierLimiter, Handler("Client Tier")))

	// Info endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Rate Limiter Test Server\nEndpoints:\n/fixed\n/sliding\n/token\n/tier\n")
	})

	fmt.Printf("Starting server on port 8080...\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}

// Wrapper to match middleware.RateLimiter interface
type slidingWrapper struct {
	limiter *ratelimit.SlidingWindowRateLimter
}

func (w *slidingWrapper) Allow(key string) bool {
	allowed, err := w.limiter.Allow(context.Background(), key)
	if err != nil {
		fmt.Printf("Sliding window error: %v\n", err)
		return false
	}
	return allowed
}

func wrapSlidingLimiter(l *ratelimit.SlidingWindowRateLimter) middleware.RateLimiter {
	return &slidingWrapper{limiter: l}
}
