package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
)

type Category struct {
	ID          pgtype.UUID `db:"id"`
	Title       pgtype.Text `db:"title"`
	Description pgtype.Text `db:"description"`
	Tag         pgtype.Text `db:"tag"`
}

func CategoryToCreateDTO(c domain.Category) Category {
	return Category{
		ID:          pgtype.UUID{Bytes: c.ID, Valid: true},
		Title:       pgtype.Text{String: c.Title, Valid: true},
		Description: pgtype.Text{String: c.Description, Valid: true},
		Tag:         pgtype.Text{String: c.Tag, Valid: true},
	}
}

func CategoryToUpdateDTO(c domain.Category) Category {
	return Category{
		ID:          pgtype.UUID{Bytes: c.ID, Valid: true},
		Title:       pgtype.Text{String: c.Title, Valid: true},
		Description: pgtype.Text{String: c.Description, Valid: true},
		Tag:         pgtype.Text{String: c.Tag, Valid: true},
	}
}

func CategoryFromGetDTO(c Category) domain.Category {
	return domain.Category{
		ID:          c.ID.Bytes,
		Title:       c.Title.String,
		Description: c.Description.String,
		Tag:         c.Tag.String,
	}
}
