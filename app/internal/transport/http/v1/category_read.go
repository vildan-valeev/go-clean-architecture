package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) CategoryRead(c *fiber.Ctx) error {
	s := new(dto.CategoryReadRequest)

	if err := c.BodyParser(s); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	result, err := t.category.ReadCategory(c.Context(), *s)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).JSON(dto.CategoryReadToResponse(result))
}
