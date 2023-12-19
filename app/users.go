package app

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"
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
		HashPassword: extend.Helper.Hash(requestBody.Password),
	}
	err, rowsAffect := app.userRepo.Create(userData)
	if err != nil {
		return err
	}
	if rowsAffect == 0 {
		return &domain.UserExistErr
	}
	return nil
}
	return nil
}
