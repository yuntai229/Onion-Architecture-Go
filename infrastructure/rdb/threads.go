package rdb

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"

	"gorm.io/gorm"
)

type ThreadRepo struct {
	Db *gorm.DB
}

func NewThreadRepo(Db *gorm.DB) ports.ThreadRepo {
	return &ThreadRepo{Db}
}

func (repo *ThreadRepo) Create(threadData entity.Threads) *entity.ErrorMessage {
	if result := repo.Db.Create(&threadData); result.Error != nil {
		return &entity.DbConnectErr
	}

	return nil
}

func (repo *ThreadRepo) GetByUserId() *entity.ErrorMessage {
	return nil
}
