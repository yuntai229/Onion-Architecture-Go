package app

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"
)

type UserApp struct {
	userRepo ports.UserRepo
}

func NewUserApp(userRepo ports.UserRepo) ports.UserApp {
	return &UserApp{userRepo}
}

func (app *UserApp) Signup(requestBody dto.SignupRequest) *entity.ErrorMessage {
	userData := entity.Users{
		Name:         requestBody.Name,
		Email:        requestBody.Email,
		HashPassword: extend.Helper.Hash(requestBody.Password),
	}
	if err := app.userRepo.Create(userData); err != nil {
		return err
	}
	return nil
}

func (app *UserApp) Login(requestBody dto.LoginRequest) (string, *entity.ErrorMessage) {
	email := requestBody.Email
	userData, err := app.userRepo.GetByMail(email)
	if err != nil {
		return "", err
	}

	if validateResult := validatePassword(userData.HashPassword, requestBody.Password); !validateResult {
		return "", &entity.PasswordIncorrectErr
	}

	var authClaims = entity.UserAuthClaims{
		UserID: userData.ID,
	}
	jwt, jwtErr := authClaims.GenJwt()
	if jwtErr != nil {
		return "", &entity.TokenGenFail
	}

	return jwt, nil
}

func validatePassword(hashed, unHashed string) bool {
	if extend.Helper.Hash(unHashed) != hashed {
		return false
	}
	return true
}
