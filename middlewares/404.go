package middlewares

import (
    "github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
)

func init() {
	server := app.Server()
	server.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})
}
