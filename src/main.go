package main

import (
	"github.com/labstack/echo/v4"
	"madronetek.com/go-frameworks-eval/cmd/handlers"
	"madronetek.com/go-frameworks-eval/cmd/storage"
  )
  
  func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	
	 // Add this line
	 storage.InitDB()
	 //----------------

	e.Logger.Fatal(e.Start(":8080"))
  }