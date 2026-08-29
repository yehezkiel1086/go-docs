package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yehezkiel1086/go-photo-api-redis/config"
)

type Redis struct {
	client *redis.Client
}

func InitRedis(ctx context.Context, conf *config.Redis) (*Redis, error) {
	redisURL := fmt.Sprintf("redis://:%s@%s/%s", conf.Password, conf.URL, conf.DB)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	// Ping to check connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) GetDB() *redis.Client {
	return r.client
}
