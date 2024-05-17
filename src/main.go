package main

import (
	// "os"
	"fmt"
	"os"

	"log/slog"

	"gofr.dev/pkg/gofr"
)

func main() {

	// initialise gofr object
	app := gofr.New()

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello World!", nil
	})

	// Showcase environment variables handling
	// without setting the environment
	//
	slog.Info(fmt.Sprintf("THIS_ENV = %s",  os.Getenv("THIS_ENV")))
	
	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
