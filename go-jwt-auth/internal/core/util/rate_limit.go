package util

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	client  *redis.Client
	limit   int
	window  time.Duration
}

func NewRateLimiter(client *redis.Client, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, key string) (bool, int, error) {
	cacheKey := fmt.Sprintf("rate_limit:%s", key)

	pipe := rl.client.Pipeline()
	incr := pipe.Incr(ctx, cacheKey)
	pipe.Expire(ctx, cacheKey, rl.window)

	if _, err := pipe.Exec(ctx); err != nil {
		return false, 0, err
	}

	count := int(incr.Val())
	remaining := rl.limit - count
	if remaining < 0 {
		remaining = 0
	}

	return count <= rl.limit, remaining, nil
}