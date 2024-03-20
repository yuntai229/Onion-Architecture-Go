package ports

import (
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
)

type ThreadRepo interface {
	Create(requestId string, threadData model.Threads) *model.ErrorMessage
	GetByUserId(requestId string, pagination *model.Pagination, userId uint) ([]model.Threads, *model.ErrorMessage)
}

type ThreadApp interface {
	CreatePost(requestId string, requestBody dto.PostRequest, userId uint) *model.ErrorMessage
	GetPost(requestId string, pagination *model.Pagination, params dto.GetPostRequest) ([]model.Threads, *model.ErrorMessage)
}
