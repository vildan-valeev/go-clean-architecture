package rs

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func InsertItemRS(ctx context.Context, rdb repository.RedisCache, u domain.Item) error {
	return nil
}

func UpdateItem(ctx context.Context, rdb repository.RedisCache, i domain.Item) error {
	id := i.ID.String()
	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, id, "id", id)
		rdb.HSet(ctx, id, "title", i.Title)
		rdb.HSet(ctx, id, "amount", i.Amount)
		rdb.HSet(ctx, id, "category_id", i.Category.ID)

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
