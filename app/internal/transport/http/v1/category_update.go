package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) CategoryUpdate(c *fiber.Ctx) error {
	s := new(dto.CategoryUpdateRequest)

	if err := c.BodyParser(s); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	err := t.category.UpdateCategory(c.Context(), *s)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(http.StatusOK)
}
