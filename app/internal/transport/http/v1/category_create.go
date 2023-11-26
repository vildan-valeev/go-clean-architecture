package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	_ "github.com/vildan-valeev/go-clean-architecture/docs"
)

// CategoryCreate 	godoc
// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    category
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.CategoryCreateRequest
// @Failure     500 {object} dto.Response
// @Router      /category/create [post]
func (t *Transport) CategoryCreate(c *fiber.Ctx) error {
	s := new(dto.CategoryCreateRequest)

	if err := c.BodyParser(s); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	result, err := t.category.CreateCategory(c.Context(), *s)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).JSON(dto.CategoryCreateToResponse(result))
}
