package model

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

func NewUsersModel() Users {
	return Users{}
}

func (model Users) SetHashPassword(unhashPassword string) string {
	return extend.Helper.Hash(unhashPassword)
}
