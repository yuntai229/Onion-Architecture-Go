package rdb

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ThreadRepo struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func NewThreadRepo(Db *gorm.DB, logger *zap.Logger) ports.ThreadRepo {
	return &ThreadRepo{Db, logger}
}

func (repo *ThreadRepo) Create(requestId string, threadData entity.Threads) *entity.ErrorMessage {
	repo.Logger.Info("[rdb][ThreadRepo][Create] Entry", zap.String("requestId", requestId), zap.Any("threadData", threadData))

	if result := repo.Db.Create(&threadData); result.Error != nil {
		repo.Logger.Error("[rdb][ThreadRepo][Create] Fail",
			zap.String("requestId", requestId),
			zap.Error(result.Error),
		)
		return &entity.DbConnectErr
	}

	return nil
}

func (repo *ThreadRepo) GetByUserId(requestId string, pagination *entity.Pagination, userId uint) ([]entity.Threads, *entity.ErrorMessage) {
	repo.Logger.Info("[rdb][ThreadRepo][GetByUserId] Entry",
		zap.String("requestId", requestId),
		zap.Uint("userId", userId),
		zap.Any("pagination", pagination),
	)

	var data []entity.Threads
	var count int64

	result := repo.Db.Scopes(pagination.NewDbPaginationScope()).Where("user_id = ?", userId).Order(pagination.ComposeOrderQuery()).Find(&data)
	if result.Error != nil {
		repo.Logger.Error("[rdb][ThreadRepo][GetByUserId] Fail",
			zap.String("requestId", requestId),
			zap.Error(result.Error),
		)
		return data, &entity.DbConnectErr
	}
	repo.Db.Model(data).Where("user_id = ?", userId).Count(&count)
	pagination.CalculatePage(count)

	return data, nil
}
