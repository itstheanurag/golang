package ratelimit

import (
	"sync"
	"time"
)

// FixedWindowRateLimiter tracks request counts within fixed time intervals.
// Use Case: Suitable for preventing basic API abuse where a hard limit per minute/hour is sufficient.
type FixedWindowRateLimiter struct {
	// mu ensures thread-safe access to the requests map across concurrent API calls.
	mu sync.Mutex
	// requests stores the usage data for each client (keyed by IP/API-Key).
	requests map[string]*windowRequestData
	// limit is the maximum number of requests a client can make in one window.
	limit int
	// windowSize is the duration of the fixed window (e.g., 1 minute).
	windowSize time.Duration
	// cleanupInterval defines how often we scan and remove expired client windows.
	cleanupInterval time.Duration
}

// windowRequestData captures the state for a specific client's current quota.
type windowRequestData struct {
	// count is the number of requests made in the current window.
	count int
	// expiresAt is the timestamp when the current window ends.
	expiresAt time.Time
}

func NewFixedWindowRateLimiter(limit int, windowSize time.Duration) *FixedWindowRateLimiter {
	limiter := &FixedWindowRateLimiter{
		limit:           limit,
		windowSize:      windowSize,
		requests:        make(map[string]*windowRequestData),
		cleanupInterval: 10 * time.Minute, // Default interval to prevent memory bloat.
	}

	// Logic: We must eventually clear memory for clients who stop visiting.
	// This background routine ensures the server doesn't "leak" memory over days of operation.
	go limiter.cleanup()

	return limiter
}

// Allow implements fixed-window logic: O(1) time and memory per active client.
func (l *FixedWindowRateLimiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	data, exists := l.requests[key]

	// Logic: If the client is new or the previous window has expired,
	// we initialize a new window starting from 'now'.
	if !exists || now.After(data.expiresAt) {
		l.requests[key] = &windowRequestData{
			count:     1,
			expiresAt: now.Add(l.windowSize),
		}
		return true
	}

	// Logic: Within the current window, we simply increment the counter.
	if data.count < l.limit {
		data.count++
		return true
	}

	// Disallow if the counter has reached the limit for the current time block.
	return false
}

// cleanup periodically removes entries from the map that are past their expiration.
// This prevents the 'requests' map from growing indefinitely.
func (l *FixedWindowRateLimiter) cleanup() {
	ticker := time.NewTicker(l.cleanupInterval)
	for range ticker.C {
		l.mu.Lock()
		now := time.Now()
		for key, data := range l.requests {
			// If the window has expired, it's safe to remove.
			// The next time this client visits, a new window will be created.
			if now.After(data.expiresAt) {
				delete(l.requests, key)
			}
		}
		l.mu.Unlock() // Use manual unlock here because defer inside a loop only executes when the function returns.
	}
}
