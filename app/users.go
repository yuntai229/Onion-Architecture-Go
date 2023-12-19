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
	if err := app.userRepo.Create(userData); err != nil {
		return err
	}
	return nil
}

func (app *UserApp) Login(requestBody dto.LoginRequest) (string, *domain.ErrorMessage) {
	email := requestBody.Email
	userData, err := app.userRepo.GetByMail(email)
	if err != nil {
		return "", err
	}

	if validateResult := validatePassword(userData.HashPassword, requestBody.Password); !validateResult {
		return "", &domain.PasswordIncorrectErr
	}

	jwt, jwtErr := extend.Helper.GenJwt(userData.ID)
	if jwtErr != nil {
		return "", &domain.TokenGenFail
	}

	return jwt, nil
}

func validatePassword(hashed, unHashed string) bool {
	if extend.Helper.Hash(unHashed) != hashed {
		return false
	}
	return true
}
