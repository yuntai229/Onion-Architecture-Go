package domain

import "onion-architecrure-go/dto"

type UserApp interface {
	Signup(requestBody dto.SignupRequest)
}

type UserRepo interface{}

type User struct {
	Name         string
	Email        string
	HashPassword string
}
