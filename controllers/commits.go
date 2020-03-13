package controllers

import (
    "fmt"
	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
    "github.com/soonoo/committrs-server/db"
    "github.com/soonoo/committrs-server/models"
)

type CommitRequest struct {
    Hash string `json:"hash"`
    UserId uint `json:"userId"`
    RepoId uint `json:"repoId"`
}

func init() {
	server := app.Server()
    DB := db.GetDB()

	commitsRouter := server.Group("/commits")

	commitsRouter.Put("/", func(c *fiber.Ctx) {
        var commitsRequest []CommitRequest

        if err:= c.BodyParser(&commitsRequest); err != nil {
            fmt.Printf(err.Error())
            c.Status(400).Send()
            return
        }

        for _, commit := range commitsRequest {
            commitObject := models.Commit{Hash: commit.Hash, RepoId: commit.RepoId, UserId: commit.UserId}
            DB.Create(&commitObject)
        }

		c.Status(200).Send()
	})
}
