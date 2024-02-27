package handler

import (
	"net/http"
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"

	"github.com/gin-gonic/gin"
)

type ThreadHandler struct {
	threadApp ports.ThreadApp
}

func NewThreadHandler(threadApp ports.ThreadApp) *ThreadHandler {
	return &ThreadHandler{threadApp}
}

func (handler *ThreadHandler) CreatePost(ctx *gin.Context) {
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

func (handler *ThreadHandler) GetPost(ctx *gin.Context) {
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

	if params.UserId == 0 {
		params.UserId = userId.(uint)
	}

	threadData, err := handler.threadApp.GetPost(&params.Pagination, params)
	if err != nil {
		newErr := *err
		res := entity.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	var resData []dto.GetPostresponse
	for _, element := range threadData {
		resData = append(resData, dto.GetPostresponse{
			Id:        element.ID,
			UserId:    element.UserId,
			Content:   element.Content,
			CreatedAt: extend.Helper.FormatToTimeString(element.CreatedAt),
			UpdatedAt: extend.Helper.FormatToTimeString(element.UpdatedAt),
		})
	}

	res := entity.Response.ResWithSucc(params.Pagination.Format(resData))
	ctx.JSON(http.StatusOK, res)
}
