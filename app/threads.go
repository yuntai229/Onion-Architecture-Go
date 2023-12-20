package app

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type ThreadApp struct {
	threadRepo domain.ThreadRepo
}

func NewThreadApp(threadRepo domain.ThreadRepo) domain.ThreadApp {
	return &ThreadApp{threadRepo}
}

func (app *ThreadApp) CreatePost(requestBody dto.PostRequest) *domain.ErrorMessage {
	threadData := domain.Threads{
		Content: requestBody.Content,
	}
	return app.threadRepo.Create(threadData)
}
