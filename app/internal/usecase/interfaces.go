package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

type Category interface {
	CreateCategory(ctx context.Context, c dto.CategoryCreateRequest) (uuid.UUID, error)
	ReadCategory(ctx context.Context, c dto.CategoryReadRequest) (domain.Category, error)
	UpdateCategory(ctx context.Context, c dto.CategoryUpdateRequest) error
	DeleteCategory(ctx context.Context, c dto.CategoryDeleteRequest) error
}

type Item interface {
	CreateItem(ctx context.Context, i dto.ItemCreateRequest) (uuid.UUID, error)
	ReadItem(ctx context.Context, i dto.ItemReadRequest) (domain.Item, error)
	UpdateItem(ctx context.Context, i dto.ItemUpdateRequest) error
	DeleteItem(ctx context.Context, i dto.ItemDeleteRequest) error
}
