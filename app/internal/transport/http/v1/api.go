package v1

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
)

type Transport struct {
	category Category
	item     Item
}
type DI struct {
	Category Category
	Item     Item
}

func NewTransport(di DI) *Transport {
	return &Transport{
		item:     di.Item,
		category: di.Category,
	}
}

func (t *Transport) Register() *fiber.App {
	app := fiber.New()

	app.Post("/item/create", t.ItemCreate)
	app.Get("/item", t.ItemRead)
	app.Delete("/item", t.ItemDelete)
	app.Post("/item", t.ItemUpdate)
	app.Post("/category/create", t.CategoryCreate)
	app.Get("/category", t.CategoryRead)
	app.Delete("/category", t.CategoryDelete)
	app.Post("/category", t.CategoryUpdate)

	return app
}

/*
Интерфейсы от бизнес слоя - Usecase.
*/

type Category interface {
	CreateCategory(ctx context.Context, u dto.CategoryDtoRequest) (uint64, error)
	ReadCategory(ctx context.Context, id uint64) (domain.Category, error)
	UpdateCategory(ctx context.Context, u *dto.CategoryDtoRequest) error
	DeleteCategory(ctx context.Context, id uint64) error
}

type Item interface {
	CreateItem(ctx context.Context, u dto.ItemCreateDtoRequest) (uint64, error)
	ReadItem(ctx context.Context, id uint64) (domain.Item, error)
	UpdateItem(ctx context.Context, u *dto.ItemUpdateDtoRequest) (uint64, error)
	DeleteItem(ctx context.Context, id uint64) error
}
