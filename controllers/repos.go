package controllers

import (
    "fmt"
	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
    "github.com/soonoo/committrs-server/db"
    "github.com/soonoo/committrs-server/models"
)

type RepoRequest struct {
    Name string `json:"name" xml:"name" form:"name"`
    Owner string `json:"Owner" xml:"Owner" form:"Owner"`
}

func init() {
	server := app.Server()
    DB := db.GetDB()

	reposRouter := server.Group("/repos")

	reposRouter.Put("/", func(c *fiber.Ctx) {
        var repoRequest RepoRequest

        if err:= c.BodyParser(&repoRequest); err != nil {
            fmt.Printf(err.Error())
            c.Status(400).Send()
        }

        repo := models.Repo{Name: repoRequest.Name, Owner: repoRequest.Owner}
        DB.Create(&repo)
		c.JSON(repo)
	})
}
