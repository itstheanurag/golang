package ratelimit

import (
	"sync"
	"time"
)

type TierConfig struct {
	Rate     float64
	Capacity int
}

type ClientTierRateLimiter struct {
	mu          sync.Mutex
	buckets     map[string]*bucket
	tiers       map[string]TierConfig
	clientTiers map[string]string
	defaultTier TierConfig
}

func NewClientTierRateLimiter(defaultTier TierConfig) *ClientTierRateLimiter {
	return &ClientTierRateLimiter{
		buckets:     make(map[string]*bucket),
		tiers:       make(map[string]TierConfig),
		clientTiers: make(map[string]string),
		defaultTier: defaultTier,
	}
}

func (t *ClientTierRateLimiter) AddTier(name string, tier TierConfig) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tiers[name] = tier
}

func (t *ClientTierRateLimiter) SetClientTier(clientKey string, tierName string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.clientTiers[clientKey] = tierName

	delete(t.buckets, clientKey)
}

func (t *ClientTierRateLimiter) Allow(key string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	config := t.defaultTier

	if tierName, exists := t.clientTiers[key]; exists {
		if tierConfig, tierExists := t.tiers[tierName]; tierExists {
			config = tierConfig
		}
	}

	now := time.Now()
	b, exists := t.buckets[key]
	if !exists {
		t.buckets[key] = &bucket{
			tokens:        float64(config.Capacity) - 1,
			lastUpdatedAt: now,
		}
		return true
	}

	elapsed := now.Sub(b.lastUpdatedAt).Seconds()
	b.tokens += elapsed * config.Rate
	if b.tokens > float64(config.Capacity) {
		b.tokens = float64(config.Capacity)
	}
	b.lastUpdatedAt = now

	if b.tokens >= 1 {
		b.tokens -= 1
		return true
	}

	return false
}
