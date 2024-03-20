package dto

import "onion-architecrure-go/domain/model"

type PostRequest struct {
	Content string `json:"content" binding:"required"`
}

type GetPostRequest struct {
	model.Pagination
	UserId uint `form:"userId"`
}

type GetPostresponse struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"userId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
