package app

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type UserApp struct{}

func NewUserApp() domain.UserApp {
	return &UserApp{}
}

func (app *UserApp) Signup(requestBody dto.SignupRequest) *domain.ErrorMessage {
	return nil
}
