package pg

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func InsertCategory(ctx context.Context, db repository.DB, u domain.Category) (string, error) {
	var id int64

	return id, nil
}

func GetCategory(ctx context.Context, db repository.DB, id uint64) (models.Category, error) {
	return domain.Item{}, nil
}

func UpdateCategory(ctx context.Context, db repository.DB, u domain.Category) error {
	return nil
}

func DeleteCategory(ctx context.Context, db repository.DB, id uint64) error {
	return nil
}
