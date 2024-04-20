package ports

import (
	"onion-architecrure-go/domain/constant"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
)

type UserApp interface {
	Signup(requestId string, requestBody dto.SignupRequest) *constant.ErrorMessage
	Login(requestId string, requestBody dto.LoginRequest) (string, *constant.ErrorMessage)
}

type UserRepo interface {
	Create(requestId string, userData model.Users) *constant.ErrorMessage
	GetByMail(requestId string, mail string) (model.Users, *constant.ErrorMessage)
}
