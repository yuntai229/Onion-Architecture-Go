package ports

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type ThreadRepo interface {
	Create(threadData entity.Threads) *entity.ErrorMessage
	GetByUserId(pagination *entity.Pagination, userId uint) ([]entity.Threads, *entity.ErrorMessage)
}

type ThreadApp interface {
	CreatePost(requestBody dto.PostRequest, userId uint) *entity.ErrorMessage
	GetPost(pagination *entity.Pagination, params dto.GetPostRequest) ([]entity.Threads, *entity.ErrorMessage)
}
