package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
)

type CategoryRepository struct {
	pg PostgresDB
	mg MongoQuery
}

func NewCategoryRepo(pg PostgresDB, mg MongoQuery) *CategoryRepository {
	return &CategoryRepository{pg: pg, mg: mg}
}

func (r CategoryRepository) PGInsertCategory(ctx context.Context, c domain.Category) error {
	err := PGInsertCategory(ctx, r.pg, CategoryToCreateDTO(c))
	if err != nil {
		//log.Error().Msgf("ERROR: %s", err)
		return err
	}

	return nil
}

func (r CategoryRepository) PGUpdateCategory(ctx context.Context, c domain.Category) error {
	err := PGUpdateCategory(ctx, r.pg, CategoryToUpdateDTO(c))
	if err != nil {
		return err
	}

	return nil
}

func (r CategoryRepository) PGGetCategory(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	cat, err := PGGetCategory(ctx, r.pg, id)
	if err != nil {
		return domain.Category{}, err
	}

	return CategoryFromGetDTO(cat), nil
}

func (r CategoryRepository) PGDeleteCategory(ctx context.Context, id uuid.UUID) error {
	err := PGDeleteCategory(ctx, r.pg, id)
	if err != nil {
		return err
	}

	return nil
}
