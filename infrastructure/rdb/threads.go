package rdb

import (
	domain "onion-architecrure-go/domain/entity"

	"gorm.io/gorm"
)

type ThreadRepo struct {
	Db *gorm.DB
}

func NewThreadRepo(Db *gorm.DB) domain.ThreadRepo {
	return &ThreadRepo{Db}
}

func (repo *ThreadRepo) Create(threadData domain.Threads) *domain.ErrorMessage {
	if result := repo.Db.Create(&threadData); result.Error != nil {
		return &domain.DbConnectErr
	}

	return nil
}
