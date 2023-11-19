package category

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в инфре).
type Repository interface {
	InsertCategoryDB(ctx context.Context, u domain.Category) (id int64, err error)
	UpdateCategoryDB(ctx context.Context, u domain.Category) error
	GetCategoryDB(ctx context.Context, id uint64) (domain.Category, error)
	DeleteCategoryDB(ctx context.Context, id uint64) error
	InsertCategoryRS(ctx context.Context, u domain.Category) (id int64, err error)
	UpdateCategoryRS(ctx context.Context, u domain.Category) error
	GetCategoryRS(ctx context.Context, id uint64) (domain.Category, error)
	DeleteCategoryRS(ctx context.Context, id uint64) error
}

type Service struct {
	db Repository
}

func New(db Repository) *Service {
	return &Service{
		db: db,
	}
}

// Sign Создание подписи.
func (s Service) CreateCategory(ctx context.Context, u dto.CategoryDtoRequest) (uint64, error) {
	var c domain.Category

	c.Hash = encode(s.Text, s.Key)

	return c.ID, nil
}
func (s Service) ReadCategory(ctx context.Context, id uint64) (domain.Category, error) {

	return domain.Category{}, nil
}
func (s Service) UpdateCategory(ctx context.Context, u *dto.CategoryDtoRequest) error {
	var c domain.Category

	c.Hash = encode(s.Text, s.Key)

	return c.ID, nil
}
func (s Service) DeleteCategory(ctx context.Context, id uint64) error {
	if id == 0 {
		return nil
	}

	return s.db.DeleteCategoryDB(ctx, id)
}
