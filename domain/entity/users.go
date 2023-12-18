package domain

import (
	"onion-architecrure-go/dto"

	"gorm.io/gorm"
)

type UserApp interface {
	Signup(requestBody dto.SignupRequest) *ErrorMessage
}

type UserRepo interface {
	Create(userData Users) *ErrorMessage
}

type Users struct {
	gorm.Model
	Name         string
	Email        string
	HashPassword string
}
