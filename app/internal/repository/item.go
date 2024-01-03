package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
)

type ItemRepository struct {
	pg PostgresDB
	rs RedisDB
	mg MongoQuery
}

func NewItemRepo(pg PostgresDB, rs RedisDB, mg MongoQuery) *ItemRepository {
	return &ItemRepository{
		pg: pg,
		rs: rs,
		mg: mg,
	}
}

func (r ItemRepository) PGInsertItem(ctx context.Context, i domain.Item) error {
	err := RSInsertItem(ctx, r.rs, ItemToCreateDTO(i))
	if err != nil {
		return err
	}

	return nil
}

func (r ItemRepository) PGUpdateItem(ctx context.Context, i domain.Item) error {
	err := PGUpdateItem(ctx, r.pg, ItemToUpdateDTO(i))
	if err != nil {
		return err
	}

	return nil
}

func (r ItemRepository) PGGetItem(ctx context.Context, id uuid.UUID) (domain.Item, error) {
	i, err := PGGetItem(ctx, r.pg, id)
	if err != nil {
		return domain.Item{}, err
	}

	return ItemFromGetDTO(i), nil
}

func (r ItemRepository) PGDeleteItem(ctx context.Context, id uuid.UUID) error {
	err := PGDeleteItem(ctx, r.pg, id)
	if err != nil {
		return err
	}

	return nil
}

func (r ItemRepository) RSInsertItem(ctx context.Context, i domain.Item) error {
	return nil
}

func (r ItemRepository) RSGetItem(ctx context.Context, id uuid.UUID) (domain.Item, error) {

	return domain.Item{}, nil
}

func (r ItemRepository) RSUpdateItem(ctx context.Context, i domain.Item) error {
	return nil
}

func (r ItemRepository) RSDeleteItem(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r ItemRepository) MGInsertItem(ctx context.Context, i domain.Item) error {
	return nil
}

func (r ItemRepository) MGUpdateItem(ctx context.Context, i domain.Item) error {
	return nil
}
func (r ItemRepository) MGGetItem(ctx context.Context, id uuid.UUID) (domain.Item, error) {
	return domain.Item{}, nil
}
func (r ItemRepository) MGDeleteItem(ctx context.Context, id uuid.UUID) error {
	return nil
}
