package rs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"strconv"
)

func UpdateCategory(ctx context.Context, rdb repository.RedisCache, u domain.Item) error {
	id := strconv.FormatInt(u.ID, 10) // тупое решение...

	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, id, "id", u.ID)
		rdb.HSet(ctx, id, "name", u.Name)
		rdb.HSet(ctx, id, "age", u.Age)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
