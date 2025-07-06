package routes

import (
	"goprj/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	router.POST("/api/task", controllers.AddNewTask)
	router.GET("/api/task", controllers.GetTaskOfUser)
}
