package controllers

import (
	"strconv"

	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/services/interfaces"

	"github.com/gin-gonic/gin"
)

type controller struct {
	services interfaces.PersonSvcInt
}

// NewPerson - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func NewPerson(s interfaces.PersonSvcInt) controller {
	return controller{services: s}
}

func (c controller) Create(ctx *gin.Context) {
	var body models.Person
	ctx.BindJSON(&body)

	// Might need this later when authenticating and authorizing
	person := models.Person{Name: body.Name, Email: body.Email, Password: body.Password}

	c.services.Create(ctx, person)
}

func (c controller) GetAll(ctx *gin.Context) {
	c.services.GetAll(ctx)
}

func (c controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.services.GetByID(ctx, idInt)
}

func (c controller) Update(ctx *gin.Context) {
	var body models.Person
	ctx.BindJSON(&body)
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}

	// Might need this later when authenticating and authorizing
	person := models.Person{ID: idInt, Name: body.Name, Email: body.Email, Password: body.Password}
	c.services.Update(ctx, person)
}

func (c controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}
	c.services.Delete(ctx, idInt)
}
