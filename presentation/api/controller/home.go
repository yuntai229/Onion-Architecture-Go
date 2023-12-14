package controller

import (
	"github.com/gin-gonic/gin"
)

type homeController struct{}

func NewHomeController() *homeController {
	return &homeController{}
}

func (h *homeController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
