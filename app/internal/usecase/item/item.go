package item

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в слое репо).
type Repository interface {
	InsertItemDB(ctx context.Context, i domain.Item) error
	UpdateItemDB(ctx context.Context, i domain.Item) error
	GetItemDB(ctx context.Context, id uuid.UUID) (domain.Item, error)
	DeleteItemDB(ctx context.Context, id uuid.UUID) error

	InsertItemRS(ctx context.Context, i domain.Item) error
	UpdateItemRS(ctx context.Context, i domain.Item) error
	GetItemRS(ctx context.Context, id uuid.UUID) (domain.Item, error)
	DeleteItemRS(ctx context.Context, id uuid.UUID) error
}

// Service - бизнес логика.
type Service struct {
	db Repository
}

func New(db Repository) *Service {
	return &Service{
		db: db,
	}
}

// CreateItem Создание записи.
func (s Service) CreateItem(ctx context.Context, i dto.ItemCreateRequest) (uuid.UUID, error) {

	return uuid.New(), nil
}

// ReadItem Получение записи.
func (s Service) ReadItem(ctx context.Context, i dto.ItemReadRequest) (domain.Item, error) {

	return domain.Item{}, nil
}

// UpdateItem Обновление записи.
func (s Service) UpdateItem(ctx context.Context, i dto.ItemUpdateRequest) error {

	return nil
}

// DeleteItem Удаление записи.
func (s Service) DeleteItem(ctx context.Context, i dto.ItemDeleteRequest) error {

	return nil
}
