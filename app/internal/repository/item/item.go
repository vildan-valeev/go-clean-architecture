package item

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/pg"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/rs"
)

func (r Repository) InsertItemDB(ctx context.Context, u domain.Item) (id int64, err error) {
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

func (r Repository) GetItemDB(ctx context.Context, id uint64) (domain.Item, error) {
	err := pg.GetItem(ctx, r.db, id)
	if err != nil {
		return domain.Item{}, err
	}

	return domain.Item{}, nil
}

func (r Repository) DeleteItemDB(ctx context.Context, id uint64) error {
	err := pg.DeleteItem(ctx, r.db, id)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) InsertItemRS(ctx context.Context, u domain.Item) (id int64, err error) {
	id, err = rs.InsertUser(ctx, r.db, u)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r Repository) UpdateItemRS(ctx context.Context, u domain.Item) error {
	err := rs.UpdateUserInCache(ctx, r.rs, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetItemRS(ctx context.Context, id uint64) (domain.Item, error) {
	err := rs.UpdateUser(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteItemRS(ctx context.Context, id uint64) error {
	err := rs.UpdateUser(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}
