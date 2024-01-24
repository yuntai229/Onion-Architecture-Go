package handler

import (
	"net/http"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userApp ports.UserApp
}

func NewUserHandler(userApp ports.UserApp) *UserHandler {
	return &UserHandler{userApp}
}

func (handler *UserHandler) Signup(ctx *gin.Context) {
	var requestBody dto.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := entity.MissingFieldErr
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if err := handler.userApp.Signup(requestBody); err != nil {
		newErr := *err
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	res := entity.Response.ResWithSucc(nil)
	ctx.JSON(http.StatusOK, res)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := entity.MissingFieldErr
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	jwtToken, err := handler.userApp.Login(requestBody)
	if err != nil {
		newErr := *err
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	resData := dto.LoginResponse{
		Token: jwtToken,
	}
	res := entity.Response.ResWithSucc(resData)
	ctx.JSON(http.StatusOK, res)
}
