package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) ItemCreate(c *fiber.Ctx) error {
	user := new(dto.ItemCreateRequest)

	if err := c.BodyParser(user); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	id, err := t.item.CreateItem(c.Context(), *user)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).JSON(dto.ItemCreateToResponse(id))
}
