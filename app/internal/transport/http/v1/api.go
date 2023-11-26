package v1

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/vildan-valeev/go-clean-architecture/docs"
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

// NewTransport
// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /v1
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

	app.Post("/healthcheck", HealthCheck)

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
