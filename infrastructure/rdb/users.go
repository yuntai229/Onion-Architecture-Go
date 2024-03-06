package rdb

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func NewUserRepo(Db *gorm.DB, logger *zap.Logger) ports.UserRepo {
	return &UserRepo{Db, logger}
}

func (repo *UserRepo) Create(requestId string, userData entity.Users) *entity.ErrorMessage {
	repo.Logger.Info("[rdb][UserRepo][Create] Entry", zap.String("requestId", requestId), zap.Any("userData", userData))

	result := repo.Db.Where(entity.Users{Email: userData.Email}).FirstOrCreate(&userData)
	if result.Error != nil {
		repo.Logger.Error("[rdb][UserRepo][Create] Fail",
			zap.String("requestId", requestId),
			zap.Error(result.Error),
		)
		return &entity.DbConnectErr
	}
	if result.RowsAffected == 0 {
		repo.Logger.Info("[rdb][UserRepo][Create] User has existed", zap.String("requestId", requestId))
		return &entity.UserExistErr
	}
	return nil
}

func (repo *UserRepo) GetByMail(requestId string, mail string) (entity.Users, *entity.ErrorMessage) {
	repo.Logger.Info("[rdb][UserRepo][GetByMail] Entry", zap.String("requestId", requestId), zap.String("mail", mail))

	var data entity.Users
	if result := repo.Db.First(&data, "email = ?", mail); result.Error != nil {
		if result.RowsAffected == 0 {
			repo.Logger.Info("[rdb][UserRepo][Create] User has existed", zap.String("requestId", requestId))
			return data, &entity.UserNotFoundErr
		}
		repo.Logger.Error("[rdb][UserRepo][GetByMail] Fail",
			zap.String("requestId", requestId),
			zap.Error(result.Error),
		)
		return data, &entity.DbConnectErr
	}

	return data, nil
}
