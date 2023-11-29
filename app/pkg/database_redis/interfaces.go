package database_redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Redis interface {
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
}
