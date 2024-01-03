package category

import "github.com/vildan-valeev/go-clean-architecture/internal/repository"

type CategoryUseCase struct {
	categoryRepo repository.CategoryRepo
}

func New(categoryRepo repository.CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{
		categoryRepo: categoryRepo,
	}
}
