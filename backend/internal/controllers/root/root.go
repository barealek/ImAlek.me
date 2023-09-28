package root

import (
	"github.com/gofiber/fiber/v2"
)

func InitController(app *fiber.App) {
	root := app.Group("/")

	root.Get("/", index)

}

func index(c *fiber.Ctx) error {
	return c.SendString("Metadata will show here when I get to it")
}
