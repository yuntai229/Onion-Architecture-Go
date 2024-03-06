package ports

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type UserApp interface {
	Signup(requestId string, requestBody dto.SignupRequest) *entity.ErrorMessage
	Login(requestId string, requestBody dto.LoginRequest) (string, *entity.ErrorMessage)
}

type UserRepo interface {
	Create(requestId string, userData entity.Users) *entity.ErrorMessage
	GetByMail(requestId string, mail string) (entity.Users, *entity.ErrorMessage)
}
