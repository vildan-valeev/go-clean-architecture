package pg

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
)

func DeleteItem(ctx context.Context, db repository.DB, u uint64) error {
	return nil
}

func GetItem(ctx context.Context, db repository.DB, u uint64) (domain.Item, error) {
	return domain.Item{}, nil
}

func UpdateItem(ctx context.Context, db repository.DB, u domain.Item) error {
	return nil
}

func InsertItem(ctx context.Context, db repository.DB, u domain.Item) (int64, error) {
	var id int64

	tx, err := db.Begin(ctx)
	if err != nil {
		return id, err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := tx.QueryRow(ctx,
		`INSERT INTO items (name, age) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT users_pkey DO UPDATE SET name=EXCLUDED.name, age=EXCLUDED.age RETURNING id`,
		u.Name,
		u.Age,
	).Scan(&id); err != nil {
		return id, err
	}

	return id, tx.Commit(ctx)
}
