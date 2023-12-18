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
	userData := domain.Users{
		Name:         requestBody.Name,
		Email:        requestBody.Email,
		HashPassword: requestBody.Password,
	}
	if err := app.userRepo.Create(userData); err != nil {
		return err
	}
	return nil
}