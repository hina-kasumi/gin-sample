package routes

import (
	"goprj/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/api/user", controllers.RegisterUser)
	router.GET("/api/user/all", controllers.GetAllUser)
	router.GET("/api/user", controllers.GetUser)
}
