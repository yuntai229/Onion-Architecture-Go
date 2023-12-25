package app

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/domain/ports"
	"onion-architecrure-go/dto"
)

type ThreadApp struct {
	threadRepo ports.ThreadRepo
}

func NewThreadApp(threadRepo ports.ThreadRepo) ports.ThreadApp {
	return &ThreadApp{threadRepo}
}

func (app *ThreadApp) CreatePost(requestBody dto.PostRequest, userId uint) *entity.ErrorMessage {
	threadData := entity.Threads{
		UserId:  userId,
		Content: requestBody.Content,
	}
	return app.threadRepo.Create(threadData)
}

func (app *ThreadApp) GetPost(pagination entity.PageRequest, userId uint) *entity.ErrorMessage {

	return nil
}
