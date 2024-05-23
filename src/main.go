package main

import (
	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/controllers"
	"madronetek.com/go-frameworks-eval/cmd/inits"
	"madronetek.com/go-frameworks-eval/cmd/repository"
	middlewares "madronetek.com/go-frameworks-eval/cmd/services"

	// "madronetek.com/go-frameworks-eval/cmd/models"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	inits.DB.AutoMigrate(&models.Person{})

	repo := repository.NewPerson()
	svc := middlewares.NewPerson(repo)
	ctrlr := controllers.NewPerson(svc)
	r.POST("/person", ctrlr.Create)
	r.GET("/person", ctrlr.GetAll)
	r.GET("/person/:id", ctrlr.GetByID)
	r.PUT("/person/:id", ctrlr.Update)
	r.DELETE("/person/:id", ctrlr.Delete)

	r.Run()
}
