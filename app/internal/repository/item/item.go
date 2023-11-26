package item

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/infra/pg"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func (r Repository) InsertItemDB(ctx context.Context, i domain.Item) error {
	err := pg.InsertItem(ctx, r.db, models.ItemToCreateDTO(i))
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateItemDB(ctx context.Context, i domain.Item) error {
	err := pg.UpdateItem(ctx, r.db, models.ItemToUpdateDTO(i))
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetItemDB(ctx context.Context, id uuid.UUID) (domain.Item, error) {
	i, err := pg.GetItem(ctx, r.db, id)
	if err != nil {
		return domain.Item{}, err
	}

	return models.ItemFromGetDTO(i), nil
}

func (r Repository) DeleteItemDB(ctx context.Context, id uuid.UUID) error {
	err := pg.DeleteItem(ctx, r.db, id)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) InsertItemRS(ctx context.Context, i domain.Item) error {
	return nil
}

func (r Repository) GetItemRS(ctx context.Context, id uuid.UUID) (domain.Item, error) {

	return domain.Item{}, nil
}

func (r Repository) UpdateItemRS(ctx context.Context, i domain.Item) error {
	return nil
}

func (r Repository) DeleteItemRS(ctx context.Context, id uuid.UUID) error {
	return nil
}
