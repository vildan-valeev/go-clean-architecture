package rs

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
	"strconv"
)

func InsertItemRS(ctx context.Context, u domain.Item) error {
	return nil
}

func UpdateItem(ctx context.Context, rdb repository.RedisCache, u domain.Item) error {
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

func GetItemRS(ctx context.Context, id uuid.UUID) (models.Item, error) {
	return models.Item{}, nil
}
func DeleteItemRS(ctx context.Context, id uuid.UUID) error {
	return nil
}
