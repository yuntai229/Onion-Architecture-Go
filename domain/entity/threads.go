package domain

import (
	"onion-architecrure-go/dto"

	"gorm.io/gorm"
)

type Threads struct {
	gorm.Model
	UserId  uint
	Content string
	Users   Users `gorm:"foreignKey:UserId"`
}

type ThreadRepo interface {
	Create(threadData Threads) *ErrorMessage
}

type ThreadApp interface {
	CreatePost(requestBody dto.PostRequest) *ErrorMessage
}
