package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HomeHandler struct {
	Logger *zap.Logger
}

func NewHomeHandler(logger *zap.Logger) *HomeHandler {
	return &HomeHandler{logger}
}

func (h *HomeHandler) Ping(ctx *gin.Context) {
	h.Logger.Info("test")
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
