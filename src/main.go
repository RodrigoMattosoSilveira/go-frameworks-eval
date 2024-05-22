package main

import (
	"madronetek.com/go-frameworks-eval/cmd/inits"
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

 r.Run()
}