package controllers

import (
	dtos "goprj/DTOs"
	"goprj/entities"
	"goprj/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var req dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user := &entities.User{Email: req.Email, Password: req.Password}
	dbUser, err := services.FindOneUser(*user)

	if err != nil || dbUser.IsValidPassword(user.Password) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	tokenString, err := services.GenToken(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
