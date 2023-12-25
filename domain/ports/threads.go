package ports

import (
	"onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"
)

type ThreadRepo interface {
	Create(threadData entity.Threads) *entity.ErrorMessage
	//GetByUserId(pagination dto.PageRequest, userId uint) *ErrorMessage
}

type ThreadApp interface {
	CreatePost(requestBody dto.PostRequest, userId uint) *entity.ErrorMessage
	GetPost(pagination entity.PageRequest, userId uint) *entity.ErrorMessage
}
