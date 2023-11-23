package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/dto"
	"net/http"
)

func (t *Transport) ItemDelete(c *fiber.Ctx) error {
	item := new(dto.ItemDeleteRequest)

	if err := c.BodyParser(item); err != nil {
		log.Error().Msgf("Ошибка парсинга входящих данных: %v ", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	err := t.item.DeleteItem(c.Context(), *item)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).JSON(dto.ItemDeleteToResponse())
}
