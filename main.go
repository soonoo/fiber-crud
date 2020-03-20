package main

import (
    "os"
    "os/signal"
    "syscall"
    "fmt"

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

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        fmt.Print("exiting ...")
        err := server.Shutdown()
        if err != nil {
            fmt.Print(err.Error())
        }
    }()

    if err := server.Listen(8001); err != nil {
        fmt.Print(err.Error())
        return
    }
}
