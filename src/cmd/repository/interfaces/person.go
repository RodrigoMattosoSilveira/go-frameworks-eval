package interfaces

import (
	"github.com/gin-gonic/gin"
	"madronetek.com/go-frameworks-eval/cmd/models"
)

type PersonRepoInt interface {
	Create(ctx *gin.Context, person models.Person)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context, id int64)
	Update(ctx *gin.Context, person models.Person)
	Delete(ictx *gin.Context, id int64)
}
