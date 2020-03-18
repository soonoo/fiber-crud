package controllers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
	"github.com/soonoo/committrs-server/db"
	"github.com/soonoo/committrs-server/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type RepoRequest struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Owner string `json:"owner" xml:"owner" form:"owner"`
}

// CreateRepo godoc
// @Summary Create a GitHub repository
// @Description Create a GitHub repository
// @ID create-github-repo
// @Accept json
// @Produce json
// @Param repo body RepoRequest true "repository"
// @Success 200 {object} RepoRequest
// @Tags repos
// @Router /repos [put]
func createRepo(c *fiber.Ctx) {
	DB := db.GetDB()

	var repoRequest RepoRequest

	if err := c.BodyParser(&repoRequest); err != nil {
		fmt.Printf(err.Error())
		c.Status(400).Send()
		return
	}

	repo := models.Repo{Name: repoRequest.Name, Owner: repoRequest.Owner}
	err := repo.Insert(c.Fasthttp, DB, boil.Infer())
	if err != nil {
		fmt.Printf(err.Error())
		c.Status(500).Send()
		return
	}
	c.JSON(repo)
}

func init() {
	server := app.Server()
	reposRouter := server.Group("/repos")

	reposRouter.Put("/", createRepo)
}
