package main

import (
	"example.com/Test/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DB struct{}

func main() {
	app := echo.New()
	app.Use(middleware.Logger())

	homeHandler := handler.HomeHandler{}
	app.GET("/", homeHandler.HandleHomeShow)

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.GET("/clicked", userHandler.HandleUserShow)

	app.Start(":8100")
}
