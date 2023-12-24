package handler

import (
	"net/http"
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"

	"github.com/gin-gonic/gin"
)

type threadHandler struct {
	threadApp domain.ThreadApp
}

func NewThreadHandler(threadApp domain.ThreadApp) *threadHandler {
	return &threadHandler{threadApp}
}

func (handler *threadHandler) CreatePost(ctx *gin.Context) {
	var requestBody dto.PostRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := domain.MissingFieldErr
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	userId, ok := ctx.Get("UserId")
	if !ok {
		newErr := domain.TokenInvalidErr
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if err := handler.threadApp.CreatePost(requestBody, userId.(uint)); err != nil {
		newErr := *err
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	res := domain.Response.ResWithSucc(nil)
	ctx.JSON(http.StatusOK, res)
}
