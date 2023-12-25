package ports

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type UserApp interface {
	Signup(requestBody dto.SignupRequest) *entity.ErrorMessage
	Login(requestBody dto.LoginRequest) (string, *entity.ErrorMessage)
}

type UserRepo interface {
	Create(userData entity.Users) *entity.ErrorMessage
	GetByMail(mail string) (entity.Users, *entity.ErrorMessage)
}
