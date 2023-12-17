package handler

import (
	"github.com/gin-gonic/gin"
)

type homeHandler struct{}

func NewHomeController() *homeHandler {
	return &homeHandler{}
}

func (h *homeHandler) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
