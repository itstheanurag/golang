package ratelimit

import (
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	mu         sync.Mutex
	requests   map[string]*windowRequestData
	limit      int
	windowSize time.Duration
}

type windowRequestData struct {
	count     int
	expiresAt time.Time
}

func NewFixedWindowRateLimiter(limit int, windowSize time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		limit:      limit,
		windowSize: windowSize,
		requests:   make(map[string]*windowRequestData),
	}
}

func (l *FixedWindowRateLimiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	data, exists := l.requests[key]

	if !exists || now.After(data.expiresAt) {
		l.requests[key] = &windowRequestData{
			count:     1,
			expiresAt: now.Add(l.windowSize),
		}

		return true
	}

	if data.count < l.limit {
		data.count++
		return true
	}

	return false

}
