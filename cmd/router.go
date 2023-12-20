package cmd

import (
	domain "onion-architecrure-go/domain/entity"
	handler "onion-architecrure-go/presentation/api/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(userApp domain.UserApp, threadApp domain.ThreadApp) *gin.Engine {
	router := gin.Default()

	homeHandler := handler.NewHomeHandler()
	userHandler := handler.NewUserHandler(userApp)
	threadHandler := handler.NewThreadHandler(threadApp)

	router.GET("/ping", homeHandler.Ping)
	router.POST("/user/signup", userHandler.Signup)
	router.POST("/user/login", userHandler.Login)

	router.POST("threads/post", threadHandler.CreatePost)

	return router
}
