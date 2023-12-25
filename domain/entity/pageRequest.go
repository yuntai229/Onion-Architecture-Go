package entity

type PageRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
