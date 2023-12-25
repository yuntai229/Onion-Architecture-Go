package dto

import "onion-architecrure-go/domain/entity"

type PostRequest struct {
	Content string `json:"content" binding:"required"`
}

type GetPostRequest struct {
	entity.PageRequest
	UserId uint `form:"userId"`
}
