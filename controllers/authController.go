package controllers

import (
	dtos "goprj/DTOs"
	"goprj/services"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var req dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// tạo jwt
	tokenString, err := services.LoginService(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	// tạo refresh token
	refreshToken, err := services.GenRefreshToken(req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	// lấy thời gian sống token
	expStr := os.Getenv("REFRESH_EXPARIATION")
	expSecs, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		log.Println("cannot parse JWT_EXPIRATION to number")
		return
	}

	// set refresh token vao cookie
	ctx.SetCookie("refresh-token",
		refreshToken,
		int(expSecs),    // thời gian sống
		"/auth/refresh", // cookie sẽ được gửi ở refresh
		"",              //domain
		false,
		true, // http only
	)

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

func RefreshToken(ctx *gin.Context) {
	refreshCookie, err := ctx.Cookie("refresh-token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token provided"})
		return
	}

	sub, err := services.ParseSubInToken(refreshCookie, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	// tạo jwt
	tokenString, err := services.GenToken(sub)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	// tạo refresh token
	refreshToken, err := services.GenRefreshToken(sub)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	// lấy thời gian sống token
	expStr := os.Getenv("REFRESH_EXPARIATION")
	expSecs, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		log.Println("cannot parse JWT_EXPIRATION to number")
		return
	}

	// set refresh token vao cookie
	ctx.SetCookie("refresh-token",
		refreshToken,
		int(expSecs),    // thời gian sống
		"/auth/refresh", // cookie sẽ được gửi ở refresh
		"",              //domain
		false,
		true, // http only
	)

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
