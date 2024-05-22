package controllers

import (
	"fmt"

	"madronetek.com/go-frameworks-eval/cmd/inits"
	"madronetek.com/go-frameworks-eval/cmd/models"

	"github.com/gin-gonic/gin"
)

func CreatePeople(ctx *gin.Context) {

	var body struct {
		Id        int64
		Name      string
		Email     string
		Password  string
		Active    string
	}
   
	ctx.BindJSON(&body)
   
	people := models.People{Name: body.Name, Email: body.Email, Password: body.Password,  Active: "yes"}
   
	fmt.Println(people)
	result := inits.DB.Create(&people)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
	 return
	}
   
	ctx.JSON(200, gin.H{"data": people})
   
}