package ratelimit

import (
	"sync"
	"time"
)

// TierConfig encapsulates the constraints for a specific service level.
type TierConfig struct {
	// Rate is the sustained tokens-per-second refill speed.
	Rate float64
	// Capacity is the allowed burst size (max tokens in bucket).
	Capacity int
}

// ClientTierRateLimiter provides business-logic-aware limiting.
// Use Case: Multi-tenant SaaS products where "Free" and "Premium" users require different quotas.
type ClientTierRateLimiter struct {
	// mu provides thread-safety for all internal maps.
	mu sync.Mutex
	// buckets stores the current TokenBucket state for each active client.
	buckets map[string]*bucket
	// tiers defines the available configurations (e.g., "silver", "gold").
	tiers map[string]TierConfig
	// clientTiers maps specific client identifiers to a named tier.
	clientTiers map[string]string
	// defaultTier is used if a client hasn't been explicitly assigned to a tier.
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

// AddTier registers a new service level with its specific constraints.
func (t *ClientTierRateLimiter) AddTier(name string, tier TierConfig) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tiers[name] = tier
}

// SetClientTier bonds a specific client to a service level.
func (t *ClientTierRateLimiter) SetClientTier(clientKey string, tierName string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.clientTiers[clientKey] = tierName

	// Why delete?
	// When a tier changes, we must reset the bucket to immediately enforce 
	// the new tier's capacity and refill rate.
	delete(t.buckets, clientKey)
}

// Allow delegates logic to the Token Bucket algorithm but with dynamic configuration.
func (t *ClientTierRateLimiter) Allow(key string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Logic: Dynamically pick the config. 
	// Check Client -> Tier mapping, then lookup Tier -> Config. Fallback to DefaultTier.
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

	// Refill based on the specific Rate assigned to this client's tier.
	elapsed := now.Sub(b.lastUpdatedAt).Seconds()
	b.tokens += elapsed * config.Rate

	// Cap at the specific Capacity assigned to this client's tier.
	if b.tokens > float64(config.Capacity) {
		b.tokens = float64(config.Capacity)
	}
	b.lastUpdatedAt = now

	// Consume and return.
	if b.tokens >= 1 {
		b.tokens -= 1
		return true
	}

	return false
}
