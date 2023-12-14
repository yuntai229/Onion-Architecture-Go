package presentation

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/presentation/api/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(userApp domain.UserApp) *gin.Engine {
	router := gin.Default()

	homeController := controller.NewHomeController()
	userController := controller.NewUserController(userApp)

	router.GET("/ping", homeController.Ping)
	router.POST("/user/signup", userController.Signup)

	return router
}
