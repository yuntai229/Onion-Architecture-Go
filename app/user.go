package app

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type UserApp struct {
	userRepo domain.UserRepo
}

func NewUserApp(userRepo domain.UserRepo) domain.UserApp {
	return &UserApp{userRepo}
}

func (app *UserApp) Signup(requestBody dto.SignupRequest) *domain.ErrorMessage {
	return nil
}
