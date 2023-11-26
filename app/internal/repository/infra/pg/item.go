package pg

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func InsertItem(ctx context.Context, db repository.DB, u models.Item) error {

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx,
		`INSERT INTO items (id, title, amount, category) VALUES ($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT items_pkey DO UPDATE SET title=EXCLUDED.title, amount=EXCLUDED.amount`,
		u.ID,
		u.Title,
		u.Amount,
		u.CategoryID,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func GetItem(ctx context.Context, db repository.DB, id uuid.UUID) (models.Item, error) {
	return models.Item{}, nil
}

func UpdateItem(ctx context.Context, db repository.DB, i models.Item) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx,
		`UPDATE items SET "title" = $1, "amount" = $2 WHERE id = $3`,
		i.Title,
		i.Amount,
		i.ID,
	)

	return err
}

func DeleteItem(ctx context.Context, db repository.DB, id uuid.UUID) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx,
		`DELETE FROM items WHERE id = $1`,
		id,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
