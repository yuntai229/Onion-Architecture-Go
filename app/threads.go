package app

import (
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
	"onion-architecrure-go/ports"

	"go.uber.org/zap"
)

type ThreadApp struct {
	ThreadRepo ports.ThreadRepo
	Logger     *zap.Logger
}

func NewThreadApp(threadRepo ports.ThreadRepo, logger *zap.Logger) ports.ThreadApp {
	return &ThreadApp{threadRepo, logger}
}

func (app *ThreadApp) CreatePost(requestId string, requestBody dto.PostRequest, userId uint) *model.ErrorMessage {
	app.Logger.Info("[App][ThreadApp][CreatePost] Entry",
		zap.String("requestId", requestId),
		zap.Uint("userId", userId),
		zap.Any("requestBody", requestBody),
	)

	threadData := model.NewThreadsModel()
	threadData.UserId = userId
	threadData.Content = requestBody.Content

	return app.ThreadRepo.Create(requestId, threadData)
}

func (app *ThreadApp) GetPost(requestId string, pagination *model.Pagination, params dto.GetPostRequest) ([]model.Threads, *model.ErrorMessage) {
	app.Logger.Info("[App][ThreadApp][GetPost] Entry",
		zap.String("requestId", requestId),
		zap.Any("pagination", pagination),
		zap.Any("params", params),
	)

	userId := params.UserId
	return app.ThreadRepo.GetByUserId(requestId, pagination, userId)

}
