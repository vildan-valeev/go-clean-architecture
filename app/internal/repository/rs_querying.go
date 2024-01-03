package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func RSInsertItem(ctx context.Context, rdb RedisDB, u Item) error {
	return nil
}

func RSUpdateItem(ctx context.Context, rdb RedisDB, i Item) error {
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

func RSGetItem(ctx context.Context, id uuid.UUID) (Item, error) {
	return Item{}, nil
}
func RSDeleteItem(ctx context.Context, id uuid.UUID) error {
	return nil
}
