package controllers

import (
	dtos "goprj/DTOs"
	"goprj/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var req dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tokenString, err := services.LoginService(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Logout(ctx *gin.Context) {
	bearer := ctx.Request.Header.Get("Authorization") // lấy header xác thực
	if err := services.LogoutService(strings.Split(bearer, " ")[1]); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}
