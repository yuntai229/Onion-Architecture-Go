package rdb

import (
	"onion-architecrure-go/domain/constant"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/ports"

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

func (repo *UserRepo) Create(requestId string, userData model.Users) *constant.ErrorMessage {
	repo.Logger.Info("[rdb][UserRepo][Create] Entry", zap.String("requestId", requestId), zap.Any("userData", userData))

	result := repo.Db.Where(model.Users{Email: userData.Email}).FirstOrCreate(&userData)
	if result.Error != nil {
		repo.Logger.Error("[rdb][UserRepo][Create] Fail",
			zap.String("requestId", requestId),
			zap.Error(result.Error),
		)
		return &constant.DbConnectErr
	}
	if result.RowsAffected == 0 {
		repo.Logger.Info("[rdb][UserRepo][Create] User has existed", zap.String("requestId", requestId))
		return &constant.UserExistErr
	}
	return nil
}

func (repo *UserRepo) GetByMail(requestId string, mail string) (model.Users, *constant.ErrorMessage) {
	repo.Logger.Info("[rdb][UserRepo][GetByMail] Entry", zap.String("requestId", requestId), zap.String("mail", mail))

	var data model.Users
	if result := repo.Db.First(&data, "email = ?", mail); result.Error != nil {
		if result.RowsAffected == 0 {
			repo.Logger.Info("[rdb][UserRepo][Create] User has existed", zap.String("requestId", requestId))
			return data, &constant.UserNotFoundErr
		}
		repo.Logger.Error("[rdb][UserRepo][GetByMail] Fail",
			zap.String("requestId", requestId),
			zap.Error(result.Error),
		)
		return data, &constant.DbConnectErr
	}

	return data, nil
}
