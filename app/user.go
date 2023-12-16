package app

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type UserApp struct{}

func NewUser() domain.UserApp {
	return &UserApp{}
}

func (u *UserApp) Signup(requestBody dto.SignupRequest) *domain.ErrorMessage {
	return nil
}
