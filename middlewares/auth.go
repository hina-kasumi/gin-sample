package middlewares

import (
	"fmt"
	"goprj/entities"
	"goprj/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// lấy thông tin người dùng từ token
func UserLoaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.Request.Header.Get("Authorization") // lấy header xác thực
		if bearer != "" {
			jwtPart := strings.Split(bearer, " ")
			if len(jwtPart) == 2 {
				jwtEncode := jwtPart[1] // lấy token

				// decode jwt
				claims, err := services.ParseAccessToken(jwtEncode)

				if err != nil {
					println(err.Error())
					return
				}

				if services.IsTokenInBlackList(jwtEncode) {
					return
				}

				// lấy thông tin người dùng
				sub := claims["sub"].(string)
				fmt.Printf("Authenticated request for email: %s\n", sub)

				user := &entities.User{
					Email: sub,
				}
				if sub != "" {
					user, _ = services.FindOneUser(*user)
				}

				ctx.Set("currentUser", *user) //set người dùng vào context
			}
		}
	}
}

// kiểm tra là người dùng
func EnforceAuthenticatedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exists := ctx.Get("currentUser") // kiểm tra người người dùng sau các filter

		if exists && user != nil {
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
	}
}
