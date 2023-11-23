package item

import (
	"context"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository/models"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в слое репо).
type Repository interface {
	InsertItemDB(ctx context.Context, u domain.Item) (id string, err error)
	UpdateItemDB(ctx context.Context, u domain.Item) error
	GetItemDB(ctx context.Context, id uint64) (models.Item, error)
	DeleteItemDB(ctx context.Context, id uint64) error

	InsertItemRS(ctx context.Context, u domain.Item) (id string, err error)
	UpdateItemRS(ctx context.Context, u domain.Item) error
	GetItemRS(ctx context.Context, id uint64) (models.Item, error)
	DeleteItemRS(ctx context.Context, id uint64) error
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
func (s Service) CreateItem(ctx context.Context, i dto.ItemCreateRequest) (uint64, error) {

	return 0, nil
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
