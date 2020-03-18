package controllers

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
	"github.com/soonoo/committrs-server/db"
	"github.com/soonoo/committrs-server/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type CommitRequest struct {
	UserId int `json:"userId"`
	RepoId int `json:"repoId"`
}

// CreateCommit godoc
// @Summary Create a commit
// @Description Create a commit
// @ID create-github-commit
// @Accept json
// @Produce json
// @Param commit body CommitRequest true "commit"
// @Success 200 {object} CommitRequest
// @Tags commits
// @Router /commits [put]
func createCommit(c *fiber.Ctx) {
	DB := db.GetDB()

	var commitRequest CommitRequest
	if err := c.BodyParser(&commitRequest); err != nil {
		fmt.Printf(err.Error())
		c.Status(400).Send()
		return
	}

	commit := models.Commit{RepoID: commitRequest.RepoId, UserID: commitRequest.UserId}
	err := commit.Insert(c.Fasthttp, DB, boil.Infer())
	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}
	c.Status(200).Send()
}

func init() {
	server := app.Server()
	commitsRouter := server.Group("/commits")

	commitsRouter.Put("/", createCommit)
}
