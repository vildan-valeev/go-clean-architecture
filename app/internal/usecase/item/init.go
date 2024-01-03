package item

import "github.com/vildan-valeev/go-clean-architecture/internal/repository"

// Service - бизнес логика.
type ItemUseCase struct {
	itemRepo repository.ItemRepo
}

func New(itemRepo repository.ItemRepo) *ItemUseCase {
	return &ItemUseCase{
		itemRepo: itemRepo,
	}
}
