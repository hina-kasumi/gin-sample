package routes

import (
	"goprj/controllers"
	"goprj/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	task := router.Group("/auth")
	task.POST("/login", controllers.Login)
	task.POST("/logout", middlewares.EnforceAuthenticatedMiddleware(), controllers.Logout)
	task.POST("/refresh", controllers.RefreshToken)
}
