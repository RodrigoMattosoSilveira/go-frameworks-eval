package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID        int64   `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
}