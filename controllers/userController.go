package controllers

import (
	dtos "goprj/DTOs"
	"goprj/entities"
	"goprj/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func RegisterUser(c *gin.Context) {
	var req dtos.NewUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dữ liệu không hợp lệ: " + err.Error(),
		})
		return
	}

	user := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if services.NewUser(&user) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể tạo người dùng",
		})
		return
	}

	c.JSON(http.StatusCreated, dtos.UserResponse{
		Email: user.Email,
		Name:  req.Name,
	})
}

func GetAllUser(c *gin.Context) {
	users, err := services.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get list users",
		})
		return
	}

	userResponses := make([]dtos.UserResponse, len(users))
	for i, user := range users {
		deepcopier.Copy(user).To(&userResponses[i])
	}

	c.JSON(http.StatusOK, userResponses)
}

func GetUser(c *gin.Context) {
	email := c.Query("email") // Lấy email từ query string
	user, err := services.FindOneUser(entities.User{Email: email})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "can't not find user",
		})
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	var response dtos.UserResponse
	deepcopier.Copy(user).To(&response)
	c.JSON(http.StatusOK, response)
}
