package handler

import (
	"errors"
	"fmt"
	"net/http"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/extend"
	"onion-architecrure-go/ports"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ThreadHandler struct {
	ThreadApp ports.ThreadApp
	Logger    *zap.Logger
}

func NewThreadHandler(threadApp ports.ThreadApp, logger *zap.Logger) *ThreadHandler {
	return &ThreadHandler{threadApp, logger}
}

func (handler *ThreadHandler) CreatePost(ctx *gin.Context) {
	requestId := fmt.Sprintf("%v", ctx.Value("requestId"))
	var requestBody dto.PostRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := model.MissingFieldErr
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][ThreadHandler][CreatePost] Request end - ShouldBindJSON Error",
			zap.String("requestId", requestId),
			zap.Any("res", res),
			zap.Error(err),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	userId, ok := ctx.Get("UserId")
	if !ok {
		newErr := model.TokenInvalidErr
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Error("[ApiHandler][ThreadHandler][CreatePost] Request end - Jwt Error",
			zap.String("requestId", requestId),
			zap.Any("res", res),
			zap.Error(errors.New("Can not parse user id from jwt token")),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if err := handler.ThreadApp.CreatePost(requestId, requestBody, userId.(uint)); err != nil {
		newErr := *err
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][ThreadHandler][CreatePost] Request end - Fail",
			zap.String("requestId", requestId),
			zap.Any("res", res),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}
	res := model.NewResModel().ResWithSucc(nil)
	handler.Logger.Info("[ApiHandler][ThreadHandler][CreatePost] Request end - Succ",
		zap.String("requestId", requestId), zap.Any("res", res),
	)
	ctx.JSON(http.StatusOK, res)
}

func (handler *ThreadHandler) GetPost(ctx *gin.Context) {
	requestId := fmt.Sprintf("%v", ctx.Value("requestId"))
	var params dto.GetPostRequest

	if err := ctx.ShouldBind(&params); err != nil {
		newErr := model.MissingFieldErr
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Error("[ApiHandler][ThreadHandler][GetPost] Request end - ShouldBindJSON Error",
			zap.String("requestId", requestId),
			zap.Any("res", res),
			zap.Error(errors.New("Can not parse user id from jwt token")),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	userId, ok := ctx.Get("UserId")
	if !ok {
		newErr := model.TokenInvalidErr
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Error("[ApiHandler][ThreadHandler][GetPost] Request end - Jwt Error",
			zap.String("requestId", requestId),
			zap.Any("res", res),
			zap.Error(errors.New("Can not parse user id from jwt token")),
		)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	if params.UserId == 0 {
		params.UserId = userId.(uint)
	}

	threadData, err := handler.ThreadApp.GetPost(requestId, &params.Pagination, params)
	if err != nil {
		newErr := *err
		res := model.NewResModel().ResWithFail(newErr)
		handler.Logger.Info("[ApiHandler][ThreadHandler][GetPost] Request end - Fail",
			zap.String("requestId", requestId),
			zap.Any("res", res),
		)
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

	res := model.NewResModel().ResWithSucc(params.Pagination.Format(resData))
	handler.Logger.Info("[ApiHandler][ThreadHandler][CreatePost] Request end - Succ",
		zap.String("requestId", requestId), zap.Any("res", res),
	)
	ctx.JSON(http.StatusOK, res)
}
