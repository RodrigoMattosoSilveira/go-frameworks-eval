package repository

import (
	"github.com/gin-gonic/gin"

	"madronetek.com/go-frameworks-eval/cmd/inits"
	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/repository/interfaces"
)

type repository struct{}

// NewPerson is a factory function for store layer that returns a interface type, UserInt
func NewPerson() interfaces.PersonRepoInt {
	return repository{}
}

// A RepositoryInt interface method
//
// Inserts a record in the user table
func (repo repository) Create(ctx *gin.Context, person models.Person) {
	result := inits.DB.Create(&person)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

func (repo repository) GetAll(ctx *gin.Context) {
	var people []models.Person

	result := inits.DB.Find(&people)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
	return
	}

	ctx.JSON(200, gin.H{"data": people})
}

func (repo repository) GetByID(ctx *gin.Context, id int64) {
	var person models.Person

	result := inits.DB.First(&person, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": person})
}

func (repo repository) Update(ctx *gin.Context, person models.Person) {
	inits.DB.Model(&person).Updates(models.Person{Name: person.Name, Email: person.Email, Password: person.Password})

	ctx.JSON(200, gin.H{"data": person})
}
func (repo repository) Delete(ctx *gin.Context, id int64) {
	inits.DB.Delete(&models.Person{}, id)

 	ctx.JSON(200, gin.H{"data": "post has been deleted successfully"})
}
