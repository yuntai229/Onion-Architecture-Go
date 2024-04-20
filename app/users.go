package app

import (
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"
	"onion-architecrure-go/ports"

	"go.uber.org/zap"
)

type UserApp struct {
	config   *model.Config
	UserRepo ports.UserRepo
	Logger   *zap.Logger
}

func NewUserApp(config *model.Config, userRepo ports.UserRepo, logger *zap.Logger) ports.UserApp {
	return &UserApp{config, userRepo, logger}
}

func (app *UserApp) Signup(requestId string, requestBody dto.SignupRequest) *model.ErrorMessage {
	app.Logger.Info("[App][UserApp][Signup] Entry",
		zap.String("requestId", requestId),
	)

	userData := model.NewUsersModel()
	userData.Name = requestBody.Name
	userData.Email = requestBody.Email
	userData.HashPassword = userData.SetHashPassword(requestBody.Password)

	if err := app.UserRepo.Create(requestId, userData); err != nil {
		return err
	}
	return nil
}

func (app *UserApp) Login(requestId string, requestBody dto.LoginRequest) (string, *model.ErrorMessage) {
	app.Logger.Info("[App][UserApp][Login] Entry",
		zap.String("requestId", requestId),
	)

	email := requestBody.Email
	userData, err := app.UserRepo.GetByMail(requestId, email)
	if err != nil {
		return "", err
	}

	if validateResult := app.validatePassword(requestId, userData.HashPassword, requestBody.Password); !validateResult {
		return "", &model.PasswordIncorrectErr
	}

	authClaims := model.NewJwtModel()
	authClaims.UserID = userData.ID
	jwt, jwtErr := authClaims.GenJwt(app.config.AppConfig.JwtKey)
	if jwtErr != nil {
		app.Logger.Error("[App][UserApp][Login] Token gen error",
			zap.String("requestId", requestId),
			zap.Error(jwtErr),
		)
		return "", &model.TokenGenFail
	}

	return jwt, nil
}

func (app *UserApp) validatePassword(requestId, hashed, unHashed string) bool {
	app.Logger.Info("[App][UserApp][validatePassword] Entry",
		zap.String("requestId", requestId),
	)

	if extend.Helper.Hash(unHashed) != hashed {
		app.Logger.Info("[App][UserApp][validatePassword] Password not correct",
			zap.String("requestId", requestId),
		)
		return false
	}
	return true
}
