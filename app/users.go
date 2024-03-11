package app

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"

	"go.uber.org/zap"
)

type UserApp struct {
	config   *entity.Config
	UserRepo ports.UserRepo
	Logger   *zap.Logger
}

func NewUserApp(config *entity.Config, userRepo ports.UserRepo, logger *zap.Logger) ports.UserApp {
	return &UserApp{config, userRepo, logger}
}

func (app *UserApp) Signup(requestId string, requestBody dto.SignupRequest) *entity.ErrorMessage {
	app.Logger.Info("[App][UserApp][Signup] Entry",
		zap.String("requestId", requestId),
	)

	userData := entity.Users{
		Name:         requestBody.Name,
		Email:        requestBody.Email,
		HashPassword: extend.Helper.Hash(requestBody.Password),
	}
	if err := app.UserRepo.Create(requestId, userData); err != nil {
		return err
	}
	return nil
}

func (app *UserApp) Login(requestId string, requestBody dto.LoginRequest) (string, *entity.ErrorMessage) {
	app.Logger.Info("[App][UserApp][Login] Entry",
		zap.String("requestId", requestId),
	)

	email := requestBody.Email
	userData, err := app.UserRepo.GetByMail(requestId, email)
	if err != nil {
		return "", err
	}

	if validateResult := app.validatePassword(requestId, userData.HashPassword, requestBody.Password); !validateResult {
		return "", &entity.PasswordIncorrectErr
	}

	var authClaims = entity.UserAuthClaims{
		UserID: userData.ID,
	}
	jwt, jwtErr := authClaims.GenJwt(app.config.JwtConfig.Key)
	if jwtErr != nil {
		app.Logger.Error("[App][UserApp][Login] Token gen error",
			zap.String("requestId", requestId),
			zap.Error(jwtErr),
		)
		return "", &entity.TokenGenFail
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
