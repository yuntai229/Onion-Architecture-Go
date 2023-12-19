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

func (repo *UserRepo) Create(userData domain.Users) (*domain.ErrorMessage, int64) {
	result := repo.Db.Where(domain.Users{Email: userData.Email}).FirstOrCreate(&userData)
	if result.Error != nil {
		return &domain.DbConnectErr, 0
	}
	return nil, result.RowsAffected
}

	return nil
}
