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
	result := repo.Db.Where(domain.Users{Email: userData.Email}).FirstOrCreate(&userData)
	if result.Error != nil {
		return &domain.DbConnectErr
	}
	if result.RowsAffected == 0 {
		return &domain.UserExistErr
	}
	return nil
}

func (repo *UserRepo) GetByMail(mail string) (*domain.Users, *domain.ErrorMessage) {
	var data domain.Users
	if result := repo.Db.First(&data, "email = ?", mail); result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, &domain.UserNotFoundErr
		}
		return nil, &domain.DbConnectErr
	}

	return &data, nil
}
