package handler

import (
	"fmt"
	"net/http"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"

	"github.com/gin-gonic/gin"
)

type threadHandler struct {
	threadApp ports.ThreadApp
}

func NewThreadHandler(threadApp ports.ThreadApp) *threadHandler {
	return &threadHandler{threadApp}
}

func (handler *threadHandler) CreatePost(ctx *gin.Context) {
	var requestBody dto.PostRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := entity.MissingFieldErr
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	userId, ok := ctx.Get("UserId")
	if !ok {
		newErr := entity.TokenInvalidErr
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if err := handler.threadApp.CreatePost(requestBody, userId.(uint)); err != nil {
		newErr := *err
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	res := entity.Response.ResWithSucc(nil)
	ctx.JSON(http.StatusOK, res)
}

func (handler *threadHandler) GetPost(ctx *gin.Context) {
	var params dto.GetPostRequest
	if err := ctx.ShouldBind(&params); err != nil {
		newErr := entity.MissingFieldErr
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	userId, ok := ctx.Get("UserId")
	if !ok {
		newErr := entity.TokenInvalidErr
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if params.UserId != 0 {
		userId = params.UserId
	}

	fmt.Println(params.UserId)
	fmt.Println(params.Page)
	fmt.Println(params.PageSize)

	if err := handler.threadApp.GetPost(params.PageRequest, userId.(uint)); err != nil {
		newErr := *err
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	res := entity.Response.ResWithSucc(nil)
	ctx.JSON(http.StatusOK, res)
}
