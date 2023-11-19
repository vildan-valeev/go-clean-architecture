package item

import (
	"context"
	"errors"

	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

// Repository - методы для работы с БД (интерфейс реализован в инфре).
type Repository interface {
	InsertItemDB(ctx context.Context, u domain.Item) (id int64, err error)
	UpdateItemDB(ctx context.Context, u domain.Item) error
	GetItemDB(ctx context.Context, id uint64) (domain.Item, error)
	DeleteItemDB(ctx context.Context, id uint64) error

	InsertItemRS(ctx context.Context, u domain.Item) (id int64, err error)
	UpdateItemRS(ctx context.Context, u domain.Item) error
	GetItemRS(ctx context.Context, id uint64) (domain.Item, error)
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

// CreateUser Создание пользователя.
func (s Service) CreateItem(ctx context.Context, u dto.ItemCreateDtoRequest) (uint64, error) {
	if u.Name == "" {
		// TODO: создать кастомные бизнесовые ошибки
		return id, errors.New("введите Имя")
	}

	user := domain.Item{
		Name: u.Name,
		Age:  u.Age,
	}

	id, err = s.db.InsertUser(ctx, user)
	if err != nil {
		return id, err
	}

	return id, nil
}

// UpdateUser Обновление пользователя.
func (s Service) ReadItem(ctx context.Context, id uint64) (domain.Item, error) {
	incrementedAgeUser := domain.Item{
		Name: user.Key,
		Age:  user.Value + 1,
	}
	err := s.db.UpdateUserInCache(ctx, incrementedAgeUser)
	if err != nil {
		return user.Value, err
	}
	//return s.db.UpdateUser(ctx, &itemID)
	// query to redis
	return incrementedAgeUser.Age, nil
}

func (s Service) UpdateItem(ctx context.Context, u *dto.ItemUpdateDtoRequest) (uint64, error) {
	incrementedAgeUser := domain.Item{
		Name: user.Key,
		Age:  user.Value + 1,
	}
	err := s.db.GetItemDB(ctx, incrementedAgeUser)
	if err != nil {
		return user.Value, err
	}
	//return s.db.UpdateUser(ctx, &itemID)
	// query to redis
	return incrementedAgeUser.Age, nil
}

func (s Service) DeleteItem(ctx context.Context, id uint64) error {
	incrementedAgeUser := domain.Item{
		Name: user.Key,
		Age:  user.Value + 1,
	}
	err := s.db.GetItemDB(ctx, incrementedAgeUser)
	if err != nil {
		return user.Value, err
	}
	//return s.db.UpdateUser(ctx, &itemID)
	// query to redis
	return incrementedAgeUser.Age, nil
}
