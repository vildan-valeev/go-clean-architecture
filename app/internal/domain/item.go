package domain

import "github.com/google/uuid"

type Item struct {
	ID     uuid.UUID
	Title  string
	Amount uint8

	Category Category
}
