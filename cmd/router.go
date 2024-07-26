package cmd

import (
	"onion-architecrure-go/presentation/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(handlers api.Handlers, middlewares api.Middlewares) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.LogTraceMiddleware.InjectRequestId())

	router.GET("/ping", handlers.HomeHandler.Ping)
	router.POST("/users/signup", handlers.UserHandler.Signup)
	router.POST("/users/login", handlers.UserHandler.Login)

	router.POST("/threads/post", middlewares.JwtMiddelware.Auth(), handlers.ThreadHandler.CreatePost)
	router.GET("/threads/post", middlewares.JwtMiddelware.Auth(), handlers.ThreadHandler.GetPost)

	return router
}
