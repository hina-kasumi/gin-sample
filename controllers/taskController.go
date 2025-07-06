package controllers

import (
	dtos "goprj/DTOs"
	"goprj/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func GetTaskOfUser(c *gin.Context) {
	email := c.Query("email") // Lấy email từ query string

	result, err := services.GetTaskOfUser(email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "can not get tasks",
		})
		return
	}

	taskResponses := make([]dtos.TaskReponse, len(result))
	for i, task := range result {
		deepcopier.Copy(task).To(&taskResponses[i])
	}

	c.JSON(http.StatusOK, taskResponses)
}

func AddNewTask(c *gin.Context) {
	var req dtos.NewTaskRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dữ liệu không hợp lệ: " + err.Error(),
		})
		return
	}

	task, err := services.AddNewTask(req.UserEmail, req.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "can not create new task",
		})
		return
	}
	c.JSON(http.StatusOK, task)
}
