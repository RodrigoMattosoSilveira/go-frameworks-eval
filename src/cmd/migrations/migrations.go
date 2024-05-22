package main

import (
	"madronetek.com/go-frameworks-eval/cmd/inits"
	"madronetek.com/go-frameworks-eval/cmd/models"
)

func init() {
 inits.LoadEnv()
 inits.DBInit()
}

func main() {
 inits.DB.AutoMigrate(&models.Post{})
}
