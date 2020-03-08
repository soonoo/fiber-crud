package app

import (
	"github.com/gofiber/fiber"
  "github.com/gofiber/fiber/middleware"
)

var App *fiber.App

func Server() *fiber.App {
	if App == nil {
		App = fiber.New()
    App.Use(middleware.Logger())
	}
	return App
}
