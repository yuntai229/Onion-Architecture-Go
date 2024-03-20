package ports

import (
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
)

type UserApp interface {
	Signup(requestId string, requestBody dto.SignupRequest) *model.ErrorMessage
	Login(requestId string, requestBody dto.LoginRequest) (string, *model.ErrorMessage)
}

type UserRepo interface {
	Create(requestId string, userData model.Users) *model.ErrorMessage
	GetByMail(requestId string, mail string) (model.Users, *model.ErrorMessage)
}
