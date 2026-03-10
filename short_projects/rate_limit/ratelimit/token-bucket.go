package ratelimit

import (
	"sync"
	"time"
)

type bucket struct {
	tokens        float64
	lastUpdatedAt time.Time
}

type TokenBucket struct {
	mu              sync.Mutex
	buckets         map[string]*bucket
	rate            float64 // tokens per second
	capacity        int     // maximum number of tokens in the bucket
	cleanupInterval time.Duration
}

func NewTokenBucketRateLimiter(rate float64, capacity int) *TokenBucket {
	token_bucket := &TokenBucket{
		rate:            rate,
		capacity:        capacity,
		cleanupInterval: 10 * time.Minute,
		buckets:         make(map[string]*bucket),
	}

	go token_bucket.cleanup()
	return token_bucket
}

func (tb *TokenBucket) Allow(key string) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	b, exists := tb.buckets[key]

	if !exists {
		tb.buckets[key] = &bucket{
			tokens:        float64(tb.capacity) - 1,
			lastUpdatedAt: now,
		}

		return true
	}

	elapsed := now.Sub(b.lastUpdatedAt).Seconds()

	b.tokens += elapsed * tb.rate

	if b.tokens > float64(tb.capacity) {
		b.tokens = float64(tb.capacity)
	}

	b.lastUpdatedAt = now

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
		defer tb.mu.Lock()
		cutoff := time.Now().Add(-tb.cleanupInterval)

		for key, b := range tb.buckets {
			if b.lastUpdatedAt.Before(cutoff) {
				delete(tb.buckets, key)
			}
		}
	}
}
