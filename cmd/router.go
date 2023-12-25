package cmd

import (
	"onion-architecrure-go/domain/ports"
	handler "onion-architecrure-go/presentation/api/handler"
	"onion-architecrure-go/presentation/api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(userApp ports.UserApp, threadApp ports.ThreadApp) *gin.Engine {
	router := gin.Default()

	jwtMiddelware := middleware.NewJwtMiddleware()

	homeHandler := handler.NewHomeHandler()
	userHandler := handler.NewUserHandler(userApp)
	threadHandler := handler.NewThreadHandler(threadApp)

	router.GET("/ping", homeHandler.Ping)
	router.POST("/user/signup", userHandler.Signup)
	router.POST("/user/login", userHandler.Login)

	router.POST("threads/post", jwtMiddelware.Auth(), threadHandler.CreatePost)
	router.GET("threads/post", jwtMiddelware.Auth(), threadHandler.GetPost)

	return router
}
