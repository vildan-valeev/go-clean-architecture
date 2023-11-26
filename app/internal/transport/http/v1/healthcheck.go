package v1

import (
	"github.com/gofiber/fiber/v2"
	"log"

	_ "github.com/vildan-valeev/go-clean-architecture/docs"
)

// HealthCheck 	godoc
// @Summary 	Show the status of server.
// @Description Проверка сервера.
// @Tags 		root
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} 			map[string]interface{}
// @Router 		/healthcheck 	[get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"Status": "Server is up and ready to work!",
	}
	log.Println("CHECK API")
	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
