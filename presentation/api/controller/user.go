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

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newErr := domain.MissingFieldErr
		res := domain.Response.ResWithFail(newErr)
		ctx.JSON(newErr.HttpCode, res)
		return
	}

	user.userApp.Signup(requestBody)
	res := domain.Response.ResWithSucc(nil)
	ctx.JSON(http.StatusOK, res)

}
