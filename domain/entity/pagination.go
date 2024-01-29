package entity

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Page     int    `form:"page,default=1" json:"page"`
	PageSize int    `form:"pageSize,default=20" json:"pageSize"`
	Total    int    `form:"total" json:"total"`
	LastPage int    `form:"lastPage" json:"lastPage"`
	OrderBy  string `form:"orderBy,default=id" json:"orderBy"`
	Sort     string `form:"sort,default=desc" json:"sort"`
}

type PageResponse struct {
	Meta       Pagination `json:"meta"`
	Collection any        `json:"collection"`
}

func (entity *Pagination) NewDbPaginationScope(table any, Db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (entity.Page - 1) * entity.PageSize
		return db.Offset(offset).Limit(entity.PageSize)
	}
}

func (entity *Pagination) ComposeOrderQuery() string {
	return fmt.Sprintf("%v %v", entity.OrderBy, entity.Sort)
}

func (entity *Pagination) CalculatePage(count int64) {
	entity.Total = int(count)
	entity.LastPage = int(math.Ceil(float64(entity.Total) / float64(entity.PageSize)))
}

func (entity *Pagination) Format(data any) PageResponse {
	return PageResponse{
		Meta: Pagination{
			Page:     entity.Page,
			PageSize: entity.PageSize,
			Total:    entity.Total,
			LastPage: entity.LastPage,
			OrderBy:  entity.OrderBy,
			Sort:     entity.Sort,
		},
		Collection: data,
	}
}
