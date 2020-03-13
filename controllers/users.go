package controllers

import (
    "fmt"
    "github.com/gofiber/fiber"
    "github.com/soonoo/committrs-server/app"
    "github.com/soonoo/committrs-server/db"
    "github.com/soonoo/committrs-server/models"
)

type UserRequest struct {
    Name string `json:"name" xml:"name" form:"name"`
}

type UserRepoRequest struct {
    Name string `json:"name" xml:"name" form:"name"`
    Owner string `json:"Owner" xml:"Owner" form:"Owner"`
}

func init() {
    server := app.Server()
    DB := db.GetDB()
    userRouter := server.Group("/users")

    userRouter.Get("/:userId", func(c *fiber.Ctx) {
        userId := c.Params("userid")
        var user models.User
        DB.Preload("Repos").Where("id = ?", userId).First(&user)
        c.JSON(user)
    })

    userRouter.Get("/", func(c *fiber.Ctx) {
        var users []models.User
        DB.Preload("Repos").Find(&users)
        c.JSON(users)
    })

    userRouter.Post("/", func(c *fiber.Ctx) {
        var user UserRequest
        if err := c.BodyParser(&user); err != nil {
            fmt.Printf(err.Error())
            c.Status(400).Send()
        }

        DB.Create(&models.User{Name: user.Name})
        c.Send(user)
    })

    userRouter.Post("/:userId/repos", func(c *fiber.Ctx) {
        userId := c.Params("userid")
        var user models.User

        var repo UserRepoRequest
        if err:= c.BodyParser(&repo); err != nil {
            fmt.Printf(err.Error())
            c.Status(400).Send()
        }

       DB.Where("id = ?", userId).
           First(&user).
           Association("Repos").
           Append(models.Repo{Owner: repo.Owner, Name: repo.Name})
    })
}
