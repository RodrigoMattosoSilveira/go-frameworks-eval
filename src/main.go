package main

import (
	"github.com/labstack/echo/v4"
	"madronetek.com/go-frameworks-eval/cmd/handlers"
	"madronetek.com/go-frameworks-eval/cmd/storage"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)

	storage.InitDB()

	// Routes
	//
	e.POST("/users", handlers.CreateUserHandler)
	e.PUT("/users/:id", handlers.UpdateUserHandler)
	e.POST("/measurements", handlers.CreateMeasurementsHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
