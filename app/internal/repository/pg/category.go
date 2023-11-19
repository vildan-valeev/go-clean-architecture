package pg

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
)

func UpdateCategory(ctx context.Context, db repository.DB, u domain.Category) error {
	return nil
}

func InsertCategory(ctx context.Context, db repository.DB, u domain.Category) (int64, error) {
	var id int64

	return id, nil
}

func DeleteCategory(ctx context.Context, db repository.DB, id uint64) error {
	return nil
}

func GetCategory(ctx context.Context, db repository.DB, id uint64) (domain.Item, error) {
	return domain.Item{}, nil
}
