package presentation

import (
	domain "onion-architecrure-go/domain/entity"
	handler "onion-architecrure-go/presentation/api/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(userApp domain.UserApp) *gin.Engine {
	router := gin.Default()

	homeController := handler.NewHomeController()
	userController := handler.NewUserController(userApp)

	router.GET("/ping", homeController.Ping)
	router.POST("/user/signup", userController.Signup)

	return router
}
