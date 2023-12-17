package rdb

import (
	domain "onion-architecrure-go/domain/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	Conn *gorm.DB
}

func NewUserRepo(conn *gorm.DB) domain.UserRepo {
	return &UserRepo{conn}
}
