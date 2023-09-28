package health

import (
	"github.com/gofiber/fiber/v2"
)

func InitController(app *fiber.App) {
	health := app.Group("/health")

	health.Get("/", index)

}

func index(c *fiber.Ctx) error {
	return c.SendString("Health is good <3")
}
