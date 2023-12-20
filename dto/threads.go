package dto

type PostRequest struct {
	Content string `json:"content" binding:"required"`
}
