package main

import (
	"github.com/soonoo/committrs-server/app"
	_ "github.com/soonoo/committrs-server/controllers"
	_ "github.com/soonoo/committrs-server/models"
)

// @title Swagger Example API
// @version 1.0
// @description https://committrs.io
// @BasePath /

// @title Swagger Example API
// @version 1.0
// @description https://commits.io
// @host localhost:8000
// @BasePath /
func main() {
	server := app.Server()
	server.Static("/swagger", "./docs")
	server.Listen(8000)
}
