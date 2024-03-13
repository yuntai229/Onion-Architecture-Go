package entity

import (
	"onion-architecrure-go/extend"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name         string
	Email        string
	HashPassword string
}

func NewUsersEntity() Users {
	return Users{}
}

func (entity Users) SetHashPassword(unhashPassword string) string {
	return extend.Helper.Hash(unhashPassword)
}
