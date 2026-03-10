package ratelimit

import (
	"sync"
	"time"
)

// bucket internal state tracks tokens for a single client.
type bucket struct {
	// tokens is the current number of available tokens (can be fractional).
	tokens float64
	// lastUpdatedAt is the timestamp of the last token consumption/refill check.
	lastUpdatedAt time.Time
}

// TokenBucket implements a standard algorithm that allows for bursts.
// Use Case: Ideal for APIs where you want to enforce a sustained average rate
// while allowing users to spike up to a maximum 'capacity' after being idle.
type TokenBucket struct {
	// mu provides thread-safety for the in-memory buckets map.
	mu sync.Mutex
	// buckets maps client keys to their specific token counts.
	buckets map[string]*bucket
	// rate is the number of tokens refilled per second.
	rate float64
	// capacity is the maximum number of tokens a bucket can hold (the burst limit).
	capacity int
	// cleanupInterval is how often we prune idle buckets from memory.
	cleanupInterval time.Duration
}

func NewTokenBucketRateLimiter(rate float64, capacity int) *TokenBucket {
	token_bucket := &TokenBucket{
		rate:            rate,
		capacity:        capacity,
		cleanupInterval: 10 * time.Minute,
		buckets:         make(map[string]*bucket),
	}

	// Why a background routine?
	// To prevent memory leaks by removing clients that haven't visited for a long time.
	go token_bucket.cleanup()
	return token_bucket
}

// Allow determines if a request is permitted based on current token availability.
func (tb *TokenBucket) Allow(key string) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	b, exists := tb.buckets[key]

	if !exists {
		// New bucket starts at full capacity minus the current request's cost.
		tb.buckets[key] = &bucket{
			tokens:        float64(tb.capacity) - 1,
			lastUpdatedAt: now,
		}
		return true
	}

	// Refill Math: Update token count based on how much time passed since the last request.
	elapsed := now.Sub(b.lastUpdatedAt).Seconds()
	b.tokens += elapsed * tb.rate

	// Ensure we don't exceed the defined burst capacity.
	if b.tokens > float64(tb.capacity) {
		b.tokens = float64(tb.capacity)
	}

	b.lastUpdatedAt = now

	// Consume one token if available.
	if b.tokens >= 1 {
		b.tokens -= 1
		return true
	}

	return false
}

func (tb *TokenBucket) cleanup() {
	ticker := time.NewTicker(tb.cleanupInterval)
	for range ticker.C {
		tb.mu.Lock()

		cutoff := time.Now().Add(-tb.cleanupInterval)

		for key, b := range tb.buckets {
			// Prune if no activity was detected within the last cleanup interval.
			if b.lastUpdatedAt.Before(cutoff) {
				delete(tb.buckets, key)
			}
		}
		tb.mu.Unlock()
	}
}
