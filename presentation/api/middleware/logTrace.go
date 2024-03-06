package middleware

import (
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type LogTraceMiddleware struct {
	Logger *zap.Logger
}

func NewLogTraceMiddleware(logger *zap.Logger) *LogTraceMiddleware {
	return &LogTraceMiddleware{logger}
}

func (middleware *LogTraceMiddleware) InjectRequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId, _ := gonanoid.New()
		ctx.Header("X-Request-Id", requestId)
		ctx.Set("requestId", requestId)
		middleware.Logger.Info("Request Start", zap.String("requestId", requestId))
		ctx.Next()
	}
}
