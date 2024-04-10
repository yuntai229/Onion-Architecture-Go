package handler

import (
	"fmt"
	"net/http"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/ports"

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
	requestId := fmt.Sprintf("%v", ctx.Value("requestId"))
	var requestBody dto.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := model.MissingFieldErr
		res := model.NewResModel().ResWithFail(newErr)
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
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][UserHandler][Signup] Request end - Fail",
			zap.String("requestId", requestId), zap.Any("res", res),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	res := model.NewResModel().ResWithSucc(nil)
	handler.Logger.Info("[ApiHandler][UserHandler][Signup] Request end - Succ",
		zap.String("requestId", requestId), zap.Any("res", res),
	)
	ctx.JSON(http.StatusOK, res)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	requestId := fmt.Sprintf("%v", ctx.Value("requestId"))
	var requestBody dto.LoginRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := model.MissingFieldErr
		res := model.NewResModel().ResWithFail(newErr)
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
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][UserHandler][Login] Request end - Fail",
			zap.String("requestId", requestId), zap.Any("res", res),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	resData := dto.LoginResponse{
		Token: jwtToken,
	}

	res := model.NewResModel().ResWithSucc(resData)
	handler.Logger.Info("[ApiHandler][UserHandler][Login] Request end - Succ",
		zap.String("requestId", requestId), zap.Any("res", res),
	)
	ctx.JSON(http.StatusOK, res)
}
