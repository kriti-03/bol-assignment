package main

import (
	"github.com/pablocrivella/mancala/internal/datastore"
	"github.com/pablocrivella/mancala/internal/engine"
	"github.com/pablocrivella/mancala/internal/service"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	gameRepo := datastore.NewGameRepo(engine.Game{})

	e := echo.New()
	e.File("/", "website/public/index.html")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	r := service.Resources{
		Games: service.GamesResource{
			GamesService: service.NewService(gameRepo),
		},
	}
	v1 := e.Group("/v1")
	v1.GET("/games/:id", r.Games.Show)
	v1.POST("/games", r.Games.Create)
	v1.PATCH("/games/:id", r.Games.Update)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
