package models

import (
	"time"

	"gorm.io/gorm"
)

type People struct {
	gorm.Model
	Id        int64
	Name      string
	Email     string
	Password  string
	Active    string
	CreatedAt time.Time
	UpdatedAt time.Time
}