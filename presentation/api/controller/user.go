package controller

import (
	"net/http"
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userApp domain.UserApp
}

func NewUserController(userApp domain.UserApp) *UserController {
	return &UserController{userApp}
}

func (user *UserController) Signup(ctx *gin.Context) {
	var requestBody dto.SignupRequest

	ctx.BindJSON(&requestBody)
	user.userApp.Signup(requestBody)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "0000",
		"message": 0,
		"data":    gin.H{},
	})

}
