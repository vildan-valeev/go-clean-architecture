package user

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/pg"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/rs"
)

func (r Repository) InsertCategoryDB(ctx context.Context, u domain.Category) (id int64, err error) {
	id, err = pg.InsertCategory(ctx, r.db, u)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r Repository) UpdateCategoryDB(ctx context.Context, u domain.Category) error {
	err := pg.UpdateCategory(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetCategoryDB(ctx context.Context, id uint64) (domain.Category, error) {
	err := pg.GetCategory(ctx, r.db, id)
	if err != nil {
		return domain.Item{}, err
	}

	return domain.Item{}, nil
}

func (r Repository) DeleteCategoryDB(ctx context.Context, id uint64) error {
	err := pg.DeleteCategory(ctx, r.db, id)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) InsertCategoryRS(ctx context.Context, u domain.Category) (id int64, err error) {
	id, err = rs.InsertCategory(ctx, r.db, u)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r Repository) UpdateCategoryRS(ctx context.Context, u domain.Category) error {
	err := rs.UpdateCategory(ctx, r.rs, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetCategoryRS(ctx context.Context, id uint64) (domain.Category, error) {
	err := rs.UpdateCategory(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteCategoryRS(ctx context.Context, id uint64) error {
	err := rs.UpdateCategory(ctx, r.db, u)
	if err != nil {
		return err
	}

	return nil
}
