package main

import (
  "github.com/soonoo/committrs-server/app"
  _ "github.com/soonoo/committrs-server/controllers"
  _ "github.com/soonoo/committrs-server/models"
)

func main() {
	server := app.Server()
	server.Listen(8000)
}
