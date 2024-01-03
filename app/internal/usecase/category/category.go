package category

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// CreateCategory Создание категории.
func (uc *CategoryUseCase) CreateCategory(ctx context.Context, c dto.CategoryCreateRequest) (uuid.UUID, error) {
	//return uuid.Nil, apperror.New(apperror.BadRequest, apperror.ErrorInvalidID)
	id := uuid.New()
	err := uc.db.PGInsertCategory(ctx, domain.Category{
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
func (uc *CategoryUseCase) ReadCategory(ctx context.Context, c dto.CategoryReadRequest) (domain.Category, error) {
	return domain.Category{}, nil
}

// UpdateCategory Обновление категории.
func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, u dto.CategoryUpdateRequest) error {
	return nil
}

// DeleteCategory Удаление категории.
func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, id dto.CategoryDeleteRequest) error {
	return nil
}
