package entity

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name         string
	Email        string
	HashPassword string
}
