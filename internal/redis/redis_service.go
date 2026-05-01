package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrCacheMiss = errors.New("cache miss")

type RedisService struct {
	client *redisClient
}

func NewRedisService(client *redisClient) *RedisService {
	return &RedisService{client}
}

// SET
func (r *RedisService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.rdb.Set(ctx, key, data, ttl).Err()
}

// GET
func (r *RedisService) Get(ctx context.Context, key string, dest any) error {
	val, err := r.client.rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		return ErrCacheMiss
	}
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// DELETE
func (r *RedisService) Delete(ctx context.Context, key string) error {
	return r.client.rdb.Del(ctx, key).Err()
}
