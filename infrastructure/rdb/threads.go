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

func (repo *ThreadRepo) GetByUserId(pagination *entity.Pagination, userId uint) ([]entity.Threads, *entity.ErrorMessage) {
	var data []entity.Threads
	var count int64

	result := repo.Db.Scopes(pagination.NewDbPaginationScope(data, repo.Db)).Where("user_id = ?", userId).Order(pagination.ComposeOrderQuery()).Find(&data)
	if result.Error != nil {
		return data, &entity.DbConnectErr
	}
	repo.Db.Model(data).Where("user_id = ?", userId).Count(&count)
	pagination.CalculatePage(count)

	return data, nil
}
