package v1

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/vildan-valeev/go-clean-architecture/docs"
	"github.com/vildan-valeev/go-clean-architecture/internal/usecase"
)

type Transport struct {
	category usecase.Category
	item     usecase.Item
}
type DI struct {
	UseCases *usecase.UseCases
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
		item:     di.UseCases.Item,
		category: di.UseCases.Category,
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
