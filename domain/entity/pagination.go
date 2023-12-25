package entity

import "gorm.io/gorm"

type Pagination PageRequest

type PageRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func (entity *PageRequest) NewDbPaginationScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (entity.Page - 1) * entity.PageSize
		return db.Offset(offset).Limit(entity.PageSize)
	}
}
