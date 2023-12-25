package entity

import (
	"gorm.io/gorm"
)

type Threads struct {
	gorm.Model
	UserId  uint
	Content string
	Users   Users `gorm:"foreignKey:UserId"`
}
