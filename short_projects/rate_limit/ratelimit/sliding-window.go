package ratelimit

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// SlidingWindowRateLimter provides high-precision rolling-window limiting using Redis.
// Use Case: Distributed systems where limits must be shared across multiple server instances.
type SlidingWindowRateLimter struct {
	// windowDuration is the rolling period (e.g., last 60s) being checked.
	windowDuration time.Duration
	// allowedPerWindow is the max number of requests allowed within any sliding window.
	allowedPerWindow int
	// redisClient manages the connection to the shared Redis instance.
	redisClient *redis.Client
	// keyPrefix helps namespace the rate limit keys in Redis.
	keyPrefix string
}

func NewSlidingWindowRateLimiter(client *redis.Client, limit int, windowSize time.Duration) *SlidingWindowRateLimter {
	return &SlidingWindowRateLimter{
		windowDuration:   windowSize,
		allowedPerWindow: limit,
		redisClient:      client,
		keyPrefix:        "ratelimit:",
	}
}

// Allow uses a Redis Sorted Set (ZSET) where each member is a request timestamp.
func (rl *SlidingWindowRateLimter) Allow(ctx context.Context, key string) (bool, error) {
	now := time.Now()
	// The window slides forward by only considering logs after windowStart.
	windowStart := now.Add(-rl.windowDuration)
	redis_key := rl.keyPrefix + key

	// Why Redis Pipelining?
	// We need to perform multiple operations (Evict old logs, Count current logs).
	// Pipelining batches these into one network request, reducing latency overhead.
	pipe := rl.redisClient.Pipeline()

	// Step 1: Remove all timestamps older than the current window boundary.
	pipe.ZRemRangeByScore(ctx, redis_key, "0", floatToString(float64(windowStart.UnixMicro())))

	// Step 2: Get the count of remaining active timestamps.
	countCmd := pipe.ZCard(ctx, redis_key)

	// Execute batch operations at once.
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	count := countCmd.Val()

	// Step 3: Check against the limit.
	if count >= int64(rl.allowedPerWindow) {
		return false, nil
	}

	// Step 4: Add the current request's timestamp (in microseconds) to the log.
	// Using UnixMicro as score/member allows sorting and handles simultaneous requests.
	member := float64(now.UnixMicro())
	err = rl.redisClient.ZAdd(ctx, redis_key, redis.Z{
		Score:  member,
		Member: member,
	}).Err()

	if err != nil {
		return false, err
	}

	// Maintenance: Ensure the key is auto-deleted if the client is inactive for a while.
	rl.redisClient.Expire(ctx, redis_key, rl.windowDuration*2)

	return true, nil
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
