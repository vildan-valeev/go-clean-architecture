package item

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// CreateItem Создание записи.
func (uc ItemUseCase) CreateItem(ctx context.Context, i dto.ItemCreateRequest) (uuid.UUID, error) {

	return uuid.New(), nil
}

// ReadItem Получение записи.
func (uc ItemUseCase) ReadItem(ctx context.Context, i dto.ItemReadRequest) (domain.Item, error) {

	return domain.Item{}, nil
}

// UpdateItem Обновление записи.
func (uc ItemUseCase) UpdateItem(ctx context.Context, i dto.ItemUpdateRequest) error {

	return nil
}

// DeleteItem Удаление записи.
func (uc ItemUseCase) DeleteItem(ctx context.Context, i dto.ItemDeleteRequest) error {

	return nil
}
