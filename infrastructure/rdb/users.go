package rdb

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo(Db *gorm.DB) ports.UserRepo {
	return &UserRepo{Db}
}

func (repo *UserRepo) Create(userData entity.Users) *entity.ErrorMessage {
	result := repo.Db.Where(entity.Users{Email: userData.Email}).FirstOrCreate(&userData)
	if result.Error != nil {
		return &entity.DbConnectErr
	}
	if result.RowsAffected == 0 {
		return &entity.UserExistErr
	}
	return nil
}

func (repo *UserRepo) GetByMail(mail string) (entity.Users, *entity.ErrorMessage) {
	var data entity.Users
	if result := repo.Db.First(&data, "email = ?", mail); result.Error != nil {
		if result.RowsAffected == 0 {
			return data, &entity.UserNotFoundErr
		}
		return data, &entity.DbConnectErr
	}

	return data, nil
}
