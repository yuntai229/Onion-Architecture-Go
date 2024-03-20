package model

import (
	"gorm.io/gorm"
)

type Threads struct {
	gorm.Model `mapstructure:",squash"`
	UserId     uint
	Content    string
	Users      Users `gorm:"foreignKey:UserId"`
}

func NewThreadsModel() Threads {
	return Threads{}
}
