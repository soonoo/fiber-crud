package controllers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/soonoo/committrs-server/app"
	"github.com/soonoo/committrs-server/db"
	"github.com/soonoo/committrs-server/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strconv"
	// "github.com/volatiletech/sqlboiler/queries/qm"
)

type UserRequest struct {
	GithubLogin string `json:"githubLogin"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatarUrl"`
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user
// @ID get-user
// @Accept json
// @Produce json
// @Success 200
// @Tags users
// @Router /users [get]
func getUsers(c *fiber.Ctx) {
	DB := db.GetDB()
	users, err := models.Users().All(c.Fasthttp, DB)
	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}

	c.JSON(users)
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @ID create-user
// @Accept json
// @Produce json
// @Param commit body UserRequest true "user"
// @Success 200
// @Tags users
// @Router /users [put]
func createUser(c *fiber.Ctx) {
	DB := db.GetDB()
	var userRequest struct {
		GithubLogin string `json:"githubLogin"`
		Email       string `json:"email"`
		AvatarURL   string `json:"avatarUrl"`
	}

	if err := c.BodyParser(&userRequest); err != nil {
		fmt.Print(err.Error())
		c.Status(400).Send()
		return
	}

	user := models.User{
		GithubLogin: userRequest.GithubLogin,
		AvatarURL:   userRequest.AvatarURL,
		Email:       userRequest.Email}

	err := user.Insert(c.Fasthttp, DB, boil.Infer())
	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}

	c.Status(200).Send()
}

// GetCommitsOfUsers godoc
// @Summary Get commits of a user
// @Description Get commits of a user
// @ID get-commits-of-user
// @Accept json
// @Produce json
// @Param userId path int true "user id"
// @Param repoId path int true "repo id"
// @Success 200
// @Tags users
// @Router /users/{userId}/commits/{repoId} [get]
func getCommitsOfUsers(c *fiber.Ctx) {
	DB := db.GetDB()
	userId := c.Params("userid")
	repoId := c.Params("repoid")

	type CommitAndRepos []struct {
		Id     int `json:"id"`
		UserId int `json:"userId"`
		// models.Commit `boil:"commits,bind"`
		models.Repo `json:"repo" boil:",bind"`
	}

	var car CommitAndRepos
	err := models.NewQuery(
		qm.Select(
			"*",
			"commits.id as id",
			"commits.user_id as user_id",
			"repos.id as \"repo.id\"",
		),
		// qm.Select("*"),
		qm.From("commits"),
		qm.Where("commits.user_id = ?", userId),
		qm.Where("commits.repo_id = ?", repoId),
		qm.InnerJoin("repos on repos.id = commits.repo_id"),
		qm.Limit(20),
	).Bind(c.Fasthttp, DB, &car)

	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}

	c.JSON(car)
}

// GetReposOfUsers godoc
// @Summary Get repos of a user
// @Description Get repos of a user
// @ID get-repos-of-user
// @Accept json
// @Produce json
// @Param userId path int true "user id"
// @Success 200
// @Tags users
// @Router /users/{userId}/repos [put]
func getReposOfUsers(c *fiber.Ctx) {
	DB := db.GetDB()
	userId := c.Params("userid")

	type UserRepo struct {
		models.Repo `boil:",bind"`
	}

	userRepos := make([]UserRepo, 0)
	err := models.NewQuery(
		qm.Select("*"),
		qm.From("repos"),
		qm.InnerJoin("user_repos ON user_repos.user_id = ?", userId),
		qm.Where("repos.id = user_repos.repo_id"),
		qm.Limit(20),
	).Bind(c.Fasthttp, DB, &userRepos)

	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}

	c.JSON(userRepos)
}

// AddRepoToUser godoc
// @Summary Add repo to a user
// @Description Add repo to a user
// @ID add-repo-to-user
// @Accept json
// @Produce json
// @Param userId path int true "user id"
// @Param repoId path int true "repo id"
// @Success 200
// @Tags users
// @Router /users/{userId}/repos/{repoId} [put]
func addRepoToUser(c *fiber.Ctx) {
	DB := db.GetDB()
	userId := c.Params("userid")
	repoId := c.Params("repoid")

	userIdInNumber, _ := strconv.Atoi(userId)
	repoIdInNumber, _ := strconv.Atoi(repoId)
	user := models.User{ID: userIdInNumber}
	repo := models.Repo{ID: repoIdInNumber}
	err := user.AddRepos(c.Fasthttp, DB, false, &repo)
	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}

	c.Send()
}

// AddCommitToUser godoc
// @Summary Add commit to a user
// @Description Add commit to a user
// @ID add-commit-to-user
// @Accept json
// @Produce json
// @Param userId path int true "user id"
// @Param repoId path int true "repo id"
// @Success 200
// @Tags users
// @Router /users/{userId}/commits/{repoId} [put]
func addCommitToUser(c *fiber.Ctx) {
	userId, _ := strconv.Atoi(c.Params("userid"))
	repoId, _ := strconv.Atoi(c.Params("repoid"))
	DB := db.GetDB()
	commit := models.Commit{RepoID: repoId, UserID: userId}
	err := commit.Insert(c.Fasthttp, DB, boil.Infer())
	if err != nil {
		fmt.Print(err.Error())
		c.Status(500).Send()
		return
	}

	c.Send()
}

func init() {
	server := app.Server()
	usersRouter := server.Group("/users")

	usersRouter.Get("/", getUsers)
	usersRouter.Put("/", createUser)
	usersRouter.Get("/:userId/commits/:repoId", getCommitsOfUsers)
	usersRouter.Get("/:userId/repos", getReposOfUsers)
	usersRouter.Put("/:userId/repos/:repoId", addRepoToUser)
	usersRouter.Put("/:userId/commits/:repoId", addCommitToUser)
}
