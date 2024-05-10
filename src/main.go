package main

import (
	"github.com/labstack/echo/v4"
	"madronetek.com/go-frameworks-eval/cmd/handlers"
  )
  
  func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	e.Logger.Fatal(e.Start(":8080"))
  }