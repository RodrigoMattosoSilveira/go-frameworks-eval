package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.PUT("/measurements/:id", handlers.UpdateMeasurementsHandler)

	e.Use(handlers.LogRequest)

	// TODO unable to mnke it work using CORS Tester: https://cors-test.codehappy.dev/
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  }))
	  
	e.Logger.Fatal(e.Start(":8080"))
}
