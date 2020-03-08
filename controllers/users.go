package controllers

import (
  "fmt"
	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
	"github.com/soonoo/committrs-server/models"
)

type UserRequest struct {
  Name string `json:"name" xml:"name" form:"name"`
}

func init() {
	server := app.Server()
  DB := models.GetDB()

	userRouter := server.Group("/users")

	userRouter.Get("/", func(c *fiber.Ctx) {
    var users []models.User
    DB.Find(&users)
    c.JSON(users)
	})

  userRouter.Post("/", func(c *fiber.Ctx) {
    var user UserRequest
    if err := c.BodyParser(&user); err != nil {
      fmt.Printf(err.Error())
    }

    DB.Create(&models.User{Name: user.Name})
    c.Send(user)
  })
}
