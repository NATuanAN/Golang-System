package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	rdb *redis.Client
}

func NewRedis(Addr string) (*redisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: Addr, DB: 0,
	})
	if err := rdb.Ping(context.Background()).Err(); err == redis.Nil {
		return nil, fmt.Errorf("redis connect error: %w", err)
	}
	return &redisClient{rdb}, nil
}

type Cache interface {
	Get(ctx context.Context, key string, dest any) error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}
