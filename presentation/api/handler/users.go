package handler

import (
	"net/http"
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userApp domain.UserApp
}

func NewUserHandler(userApp domain.UserApp) *UserHandler {
	return &UserHandler{userApp}
}

func (handler *UserHandler) Signup(ctx *gin.Context) {
	var requestBody dto.SignupRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := domain.MissingFieldErr
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if err := handler.userApp.Signup(requestBody); err != nil {
		newErr := *err
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	res := domain.Response.ResWithSucc(nil)
	ctx.JSON(http.StatusOK, res)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := domain.MissingFieldErr
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	jwtToken, err := handler.userApp.Login(requestBody)
	if err != nil {
		newErr := *err
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
	}
	resData := dto.LoginResponse{
		Token: jwtToken,
	}
	res := domain.Response.ResWithSucc(resData)
	ctx.JSON(http.StatusOK, res)
}
