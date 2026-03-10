package ratelimit

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowRateLimter struct {
	windowDuration   time.Duration
	allowedPerWindow int
	redisClient      *redis.Client
	keyPrefix        string
}

func NewSlidingWindowRateLimiter(client *redis.Client, limit int, windowSize time.Duration) *SlidingWindowRateLimter {
	return &SlidingWindowRateLimter{
		windowDuration:   windowSize,
		allowedPerWindow: limit,
		redisClient:      client,
		keyPrefix:        "ratelimit:",
	}
}

func (rl *SlidingWindowRateLimter) Allow(ctx context.Context, key string) (bool, error) {
	now := time.Now()
	windowStart := now.Add(-rl.windowDuration)
	redis_key := rl.keyPrefix + key

	pipe := rl.redisClient.Pipeline()
	// Remove entries outside the current window
	pipe.ZRemRangeByScore(ctx, redis_key, "0", floatToString(float64(windowStart.UnixMicro())))

	// Count remaining entries in the window
	countCmd := pipe.ZCard(ctx, redis_key)

	// Execute the pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	count := countCmd.Val()

	if count >= int64(rl.allowedPerWindow) {
		return false, nil
	}

	// Add current request to the sorted set
	// Using UnixMicro as both member and score ensures uniqueness
	member := float64(now.UnixMicro())
	err = rl.redisClient.ZAdd(ctx, redis_key, redis.Z{
		Score:  member,
		Member: member,
	}).Err()

	if err != nil {
		return false, err
	}

	// Set expiration on the key to auto-cleanup inactive clients
	rl.redisClient.Expire(ctx, redis_key, rl.windowDuration*2)

	return true, nil
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
