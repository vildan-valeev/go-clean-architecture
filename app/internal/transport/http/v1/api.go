package v1

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
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
