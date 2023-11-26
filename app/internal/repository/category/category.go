package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/infra/pg"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func (r Repository) InsertCategoryDB(ctx context.Context, c domain.Category) error {
	err := pg.InsertCategory(ctx, r.db, models.CategoryToCreateDTO(c))
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateCategoryDB(ctx context.Context, c domain.Category) error {
	err := pg.UpdateCategory(ctx, r.db, models.CategoryToUpdateDTO(c))
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetCategoryDB(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	cat, err := pg.GetCategory(ctx, r.db, id)
	if err != nil {
		return domain.Category{}, err
	}

	return models.CategoryFromGetDTO(cat), nil
}

func (r Repository) DeleteCategoryDB(ctx context.Context, id uuid.UUID) error {
	err := pg.DeleteCategory(ctx, r.db, id)
	if err != nil {
		return err
	}

	return nil
}
