package dto

type SignupRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	HashPassword string `json:"password" binding:"required"`
}
