package handler

import (
	"fmt"
	"net/http"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserApp ports.UserApp
	Logger  *zap.Logger
}

func NewUserHandler(userApp ports.UserApp, logger *zap.Logger) *UserHandler {
	return &UserHandler{userApp, logger}
}

func (handler *UserHandler) Signup(ctx *gin.Context) {
	requestId := fmt.Sprintf("%v", ctx.Value("RequestId"))
	var requestBody dto.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := entity.MissingFieldErr
		res := entity.NewResEntity().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][UserHandler][Signup] Request end - ShouldBindJSON Error",
			zap.String("requestId", requestId),
			zap.Any("res", res),
			zap.Error(err),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	handler.Logger.Info("[ApiHandler][UserHandler][Signup] Entry",
		zap.String("requestId", requestId),
	)

	if err := handler.UserApp.Signup(requestId, requestBody); err != nil {
		newErr := *err
		res := entity.NewResEntity().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][UserHandler][Signup] Request end - Fail",
			zap.String("requestId", requestId), zap.Any("res", res),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	res := entity.NewResEntity().ResWithSucc(nil)
	handler.Logger.Info("[ApiHandler][UserHandler][Signup] Request end - Succ",
		zap.String("requestId", requestId), zap.Any("res", res),
	)
	ctx.JSON(http.StatusOK, res)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	requestId := fmt.Sprintf("%v", ctx.Value("requestId"))
	var requestBody dto.LoginRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := entity.MissingFieldErr
		res := entity.NewResEntity().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][UserHandler][Login] Request end - ShouldBindJSON Error",
			zap.String("requestId", requestId),
			zap.Any("res", res),
			zap.Error(err),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	handler.Logger.Info("[ApiHandler][UserHandler][Login] Entry",
		zap.String("requestId", requestId),
	)

	jwtToken, err := handler.UserApp.Login(requestId, requestBody)
	if err != nil {
		newErr := *err
		res := entity.NewResEntity().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][UserHandler][Login] Request end - Fail",
			zap.String("requestId", requestId), zap.Any("res", res),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	resData := dto.LoginResponse{
		Token: jwtToken,
	}

	res := entity.NewResEntity().ResWithSucc(resData)
	handler.Logger.Info("[ApiHandler][UserHandler][Login] Request end - Succ",
		zap.String("requestId", requestId), zap.Any("res", res),
	)
	ctx.JSON(http.StatusOK, res)
}
