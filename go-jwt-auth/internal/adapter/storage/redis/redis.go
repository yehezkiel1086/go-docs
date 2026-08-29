package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
)

type Redis struct {
	client *redis.Client
}

func New(ctx context.Context, conf *config.Redis) (*Redis, error) {
	db, err := strconv.Atoi(conf.DB)
	if err != nil {
		return nil, errors.New("invalid redis db")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       db,
		Protocol: 2,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Redis{rdb}, nil
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *Redis) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *Redis) Close() error {
	return r.client.Close()
}