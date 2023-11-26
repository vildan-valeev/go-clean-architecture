package pg

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
)

func InsertCategory(ctx context.Context, db repository.DB, u models.Category) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Exec(ctx,
		`INSERT INTO categories (id, title, description, tag) VALUES ($1, $2, $3, $4) ON CONFLICT ON CONSTRAINT categories_pkey DO UPDATE SET title=EXCLUDED.title, amount=EXCLUDED.description, tag=EXCLUDED.tag`,
		u.ID,
		u.Title,
		u.Description,
		u.Tag,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func GetCategory(ctx context.Context, db repository.DB, id uuid.UUID) (models.Category, error) {
	var category models.Category

	if err := pgxscan.Get(ctx, db, &category,
		`select id, title, description, tag from categories where id = $1`,
		id,
	); err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func UpdateCategory(ctx context.Context, db repository.DB, c models.Category) error {
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

func DeleteCategory(ctx context.Context, db repository.DB, id uuid.UUID) error {
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
