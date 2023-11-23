package item

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/infra/pg"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func (r Repository) InsertItemDB(ctx context.Context, u domain.Item) (id string, err error) {
	id, err = pg.InsertItem(ctx, r.db, u)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r Repository) UpdateItemDB(ctx context.Context, u domain.Item) error {
	err := pg.UpdateItem(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetItemDB(ctx context.Context, id uint64) (models.Item, error) {

	return domain.Item{}, nil
}

func (r Repository) DeleteItemDB(ctx context.Context, id uint64) error {
	err := pg.DeleteItem(ctx, r.db, id)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) InsertItemRS(ctx context.Context, u domain.Item) (id string, err error) {
	return id, nil
}

func (r Repository) UpdateItemRS(ctx context.Context, u domain.Item) error {
	return nil
}

func (r Repository) GetItemRS(ctx context.Context, id uint64) (models.Item, error) {

	return models.Item{}, nil
}

func (r Repository) DeleteItemRS(ctx context.Context, id uint64) error {
	return nil
}
