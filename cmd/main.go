package main

import (
	"example.com/Test/handler"
	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.GET("/clicked", userHandler.HandleUserShow)

	app.Start(":8100")
}
