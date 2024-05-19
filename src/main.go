package main

import (
	// "os"
	"fmt"
	"os"

	"log/slog"

	"gofr.dev/pkg/gofr"
	"madronetek.com/go-frameworks-eval/cmd/handler"
	"madronetek.com/go-frameworks-eval/cmd/service"
	"madronetek.com/go-frameworks-eval/cmd/repository"
)

func main() {

	// initialise gofr object
	app := gofr.New()


	// Add migrations to run
	// app.Migrate(migrations.All())

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello World!", nil
	})


	repo :=repository.New()
	svc := service.New(repo)
	hdlr := handler.New(svc)


	// AddRESTHandlers creates CRUD handles for the given entity
	// Add required routes
	app.POST("/user", hdlr.Create)
	app.GET("/user", hdlr.GetAll)
	app.GET("/user/{id}", hdlr.GetByID)
	app.PUT("/user/{id}", hdlr.Update)
	app.DELETE("/user/{id}", hdlr.Delete)
	slog.Info(fmt.Sprintf("Configured = %s CRUD handler", "user"))

	// Showcase environment variables handling
	// without setting the environment
	//
	slog.Info(fmt.Sprintf("THIS_ENV = %s",  os.Getenv("THIS_ENV")))
	slog.Info(fmt.Sprintf("DB_DIALECT = %s",  os.Getenv("DB_DIALECT")))
	
	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}