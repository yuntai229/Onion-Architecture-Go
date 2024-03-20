package model

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

func (model *Pagination) NewDbPaginationScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (model.Page - 1) * model.PageSize
		return db.Offset(offset).Limit(model.PageSize)
	}
}

func (model *Pagination) ComposeOrderQuery() string {
	return fmt.Sprintf("%v %v", model.OrderBy, model.Sort)
}

func (model *Pagination) CalculatePage(count int64) {
	model.Total = int(count)
	model.LastPage = int(math.Ceil(float64(model.Total) / float64(model.PageSize)))
}

func (model *Pagination) Format(data any) PageResponse {
	return PageResponse{
		Meta: Pagination{
			Page:     model.Page,
			PageSize: model.PageSize,
			Total:    model.Total,
			LastPage: model.LastPage,
			OrderBy:  model.OrderBy,
			Sort:     model.Sort,
		},
		Collection: data,
	}
}
