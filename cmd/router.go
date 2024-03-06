package cmd

import (
	"onion-architecrure-go/presentation/api/handler"
	"onion-architecrure-go/presentation/api/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(handlers []any, middlewares []any, logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	var homeHandler *handler.HomeHandler
	var userHandler *handler.UserHandler
	var threadHandler *handler.ThreadHandler
	for _, item := range handlers {
		switch handler := item.(type) {
		case *handler.HomeHandler:
			homeHandler = handler
		case *handler.UserHandler:
			userHandler = handler
		case *handler.ThreadHandler:
			threadHandler = handler
		}
	}

	var jwtMiddelware *middleware.JwtAuthMiddleware
	var logTraceMiddleware *middleware.LogTraceMiddleware
	for _, item := range middlewares {
		switch middleware := item.(type) {
		case *middleware.JwtAuthMiddleware:
			jwtMiddelware = middleware
		case *middleware.LogTraceMiddleware:
			logTraceMiddleware = middleware
		}
	}

	router.Use(logTraceMiddleware.InjectRequestId())

	router.GET("/ping", homeHandler.Ping)
	router.POST("/users/signup", userHandler.Signup)
	router.POST("/users/login", userHandler.Login)

	router.POST("/threads/post", jwtMiddelware.Auth(), threadHandler.CreatePost)
	router.GET("/threads/post", jwtMiddelware.Auth(), threadHandler.GetPost)

	return router
}
