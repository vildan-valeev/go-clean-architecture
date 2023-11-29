package category

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в слое репо).
type Repository interface {
	InsertCategoryDB(ctx context.Context, u domain.Category) error
	UpdateCategoryDB(ctx context.Context, u domain.Category) error
	GetCategoryDB(ctx context.Context, id uuid.UUID) (domain.Category, error)
	DeleteCategoryDB(ctx context.Context, id uuid.UUID) error
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
func (s Service) CreateCategory(ctx context.Context, c dto.CategoryCreateRequest) (uuid.UUID, error) {
	//return uuid.Nil, apperror.New(apperror.BadRequest, apperror.ErrorInvalidID)
	id := uuid.New()
	err := s.db.InsertCategoryDB(ctx, domain.Category{
		ID:          uuid.New(),
		Title:       c.Title,
		Description: c.Description,
		Tag:         c.Tag,
	})
	if err != nil {
		//return uuid.Nil, apperror.New(apperror.BadRequest, apperror.ErrorInvalidID)
		return uuid.Nil, err
	}

	return id, nil
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
