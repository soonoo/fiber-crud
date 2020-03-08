package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
)

func init() {
	server := app.Server()

	reposRouter := server.Group("/repos")
	reposRouter.Get("/", func(c *fiber.Ctx) {
		c.Send("repos controller")
	})
}
