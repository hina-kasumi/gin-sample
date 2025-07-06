package routes

import (
	"goprj/controllers"
	"goprj/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	task := router.Group("/api/task")
	task.Use(middlewares.EnforceAuthenticatedMiddleware())
	task.POST("", controllers.AddNewTask)
	task.GET("", controllers.GetTaskOfUser)
}
