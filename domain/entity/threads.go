package entity

import (
	"gorm.io/gorm"
)

type Threads struct {
	gorm.Model `mapstructure:",squash"`
	UserId     uint
	Content    string
	Users      Users `gorm:"foreignKey:UserId"`
}

func NewThreadsEntity() Threads {
	return Threads{}
}
