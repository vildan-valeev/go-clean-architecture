package usecase

import (
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	"github.com/vildan-valeev/go-clean-architecture/internal/usecase/category"
	"github.com/vildan-valeev/go-clean-architecture/internal/usecase/item"
)

type UseCases struct {
	Category Category
	Item     Item
}

type Deps struct {
	Repositories *repository.Repositories
	//Host     string
	//Admin    int64
}

func NewUseCases(deps Deps) *UseCases {
	c := category.New(deps.Repositories.Category)
	i := item.New(deps.Repositories.Item)

	return &UseCases{
		Category: c,
		Item:     i,
	}
}
