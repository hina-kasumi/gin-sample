package middlewares

import (
	"fmt"
	"goprj/entities"
	"goprj/services"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UserLoaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.Request.Header.Get("Authorization") // lấy header xác thực
		if bearer != "" {
			jwtPart := strings.Split(bearer, " ")
			if len(jwtPart) == 2 {
				jwtEncode := jwtPart[1] // lấy token

				// decode jwt
				claims, err := services.ParseToken(jwtEncode)

				expToken := claims["exp"].(float64)
				if expToken < float64(time.Now().Unix()) {
					log.Println("ERROR: token is expried")
					return
				}

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
