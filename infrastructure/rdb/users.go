package rdb

import (
	domain "onion-architecrure-go/domain/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo(Db *gorm.DB) domain.UserRepo {
	return &UserRepo{Db}
}

func (repo *UserRepo) Create(userData domain.Users) *domain.ErrorMessage {
	if result := repo.Db.Create(&userData); result.Error != nil {
		return &domain.DbConnectError
	}

	return nil
}
