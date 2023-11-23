package category

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в слое репо).
type Repository interface {
	InsertCategoryDB(ctx context.Context, u domain.Category) (id int64, err error)
	UpdateCategoryDB(ctx context.Context, u domain.Category) error
	GetCategoryDB(ctx context.Context, id uint64) (domain.Category, error)
	DeleteCategoryDB(ctx context.Context, id uint64) error
}

type Service struct {
	db Repository
}

func New(db Repository) *Service {
	return &Service{
		db: db,
	}
}

// CreateCategory Создание категории.
func (s Service) CreateCategory(ctx context.Context, c dto.CategoryCreateRequest) (uint64, error) {
	return 0, nil
}

// ReadCategory Получение категории.
func (s Service) ReadCategory(ctx context.Context, c dto.CategoryReadRequest) (domain.Category, error) {
	return domain.Category{}, nil
}

// UpdateCategory Обновление категории.
func (s Service) UpdateCategory(ctx context.Context, u dto.CategoryUpdateRequest) error {
	return nil
}

// DeleteCategory Удаление категории.
func (s Service) DeleteCategory(ctx context.Context, id dto.CategoryDeleteRequest) error {
	return nil
}
