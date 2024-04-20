package ports

import (
	"onion-architecrure-go/domain/constant"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/dto"
)

type ThreadRepo interface {
	Create(requestId string, threadData model.Threads) *constant.ErrorMessage
	GetByUserId(requestId string, pagination *model.Pagination, userId uint) ([]model.Threads, *constant.ErrorMessage)
}

type ThreadApp interface {
	CreatePost(requestId string, requestBody dto.PostRequest, userId uint) *constant.ErrorMessage
	GetPost(requestId string, pagination *model.Pagination, params dto.GetPostRequest) ([]model.Threads, *constant.ErrorMessage)
}
