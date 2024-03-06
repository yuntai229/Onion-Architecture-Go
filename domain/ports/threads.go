package ports

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type ThreadRepo interface {
	Create(requestId string, threadData entity.Threads) *entity.ErrorMessage
	GetByUserId(requestId string, pagination *entity.Pagination, userId uint) ([]entity.Threads, *entity.ErrorMessage)
}

type ThreadApp interface {
	CreatePost(requestId string, requestBody dto.PostRequest, userId uint) *entity.ErrorMessage
	GetPost(requestId string, pagination *entity.Pagination, params dto.GetPostRequest) ([]entity.Threads, *entity.ErrorMessage)
}
