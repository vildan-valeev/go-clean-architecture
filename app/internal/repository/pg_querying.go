package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

func PGInsertCategory(ctx context.Context, db PostgresDB, u Category) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx,
		`INSERT INTO categories (id, title, description) VALUES ($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT categories_pk DO UPDATE SET title=EXCLUDED.title, description=EXCLUDED.description, tag=EXCLUDED.tag`,
		u.ID,
		u.Title,
		u.Description,
		u.Tag,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func PGGetCategory(ctx context.Context, db PostgresDB, id uuid.UUID) (Category, error) {
	var category Category

	if err := pgxscan.Get(ctx, db, &category,
		`select id, title, description, tag from categories where id = $1`,
		id,
	); err != nil {
		return Category{}, err
	}

	return category, nil
}

func PGUpdateCategory(ctx context.Context, db PostgresDB, c Category) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx,
		`UPDATE categories SET "title" = $1, "description" = $2, "tag" = $3 WHERE id = $4`,
		c.Title,
		c.Description,
		c.Tag,
		c.ID,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func PGDeleteCategory(ctx context.Context, db PostgresDB, id uuid.UUID) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx,
		`DELETE FROM categories WHERE id = $1`,
		id,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func PGInsertItem(ctx context.Context, db PostgresDB, u Item) error {

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

func PGGetItem(ctx context.Context, db PostgresDB, id uuid.UUID) (Item, error) {
	return Item{}, nil
}

func PGUpdateItem(ctx context.Context, db PostgresDB, i Item) error {
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

func PGDeleteItem(ctx context.Context, db PostgresDB, id uuid.UUID) error {
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
