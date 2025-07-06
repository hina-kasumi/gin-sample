package routes

import (
	"goprj/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	task := router.Group("/auth")
	task.POST("/login", controllers.Login)
}
