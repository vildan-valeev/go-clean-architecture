package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
)

type Item struct {
	ID         pgtype.UUID `db:"id"`
	Title      pgtype.Text `db:"title"`
	Amount     pgtype.Int8 `db:"amount"`
	CategoryID pgtype.UUID `db:"category_id"`
}

func ItemToCreateDTO(item domain.Item) Item {
	return Item{
		ID:         pgtype.UUID{Bytes: item.ID, Valid: true},
		Title:      pgtype.Text{String: item.Title, Valid: true},
		Amount:     pgtype.Int8{Int64: item.Amount, Valid: true},
		CategoryID: pgtype.UUID{Bytes: item.Category.ID, Valid: true},
	}
}

func ItemToUpdateDTO(item domain.Item) Item {
	return Item{
		ID:         pgtype.UUID{Bytes: item.ID, Valid: true},
		Title:      pgtype.Text{String: item.Title, Valid: true},
		Amount:     pgtype.Int8{Int64: item.Amount, Valid: true},
		CategoryID: pgtype.UUID{Bytes: item.Category.ID, Valid: true},
	}
}

func ItemFromGetDTO(i Item) domain.Item {
	return domain.Item{
		ID:     i.ID.Bytes,
		Title:  i.Title.String,
		Amount: i.Amount.Int64,
		Category: domain.Category{
			ID: i.CategoryID.Bytes,
		},
	}
}
