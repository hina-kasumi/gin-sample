package middlewares

import (
	"fmt"
	"goprj/entities"
	"goprj/services"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserLoaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.Request.Header.Get("Authorization") // lấy header xác thực
		if bearer != "" {
			jwtPart := strings.Split(bearer, " ")
			if len(jwtPart) == 2 {
				jwtEncode := jwtPart[1] // lấy token

				// decode jwt
				token, err := jwt.Parse(jwtEncode, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method")
					}
					secret := []byte(os.Getenv("JWT_SECRET"))
					return secret, nil
				})

				if err != nil {
					println(err.Error())
					return
				}

				// lấy thông tin người dùng
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					email := claims["email"].(string)
					fmt.Printf("Authenticated request for email: %s\n", email)

					user := &entities.User{
						Email: email,
					}
					if email != "" {
						user, _ = services.FindOneUser(*user)
					}

					ctx.Set("currentUser", *user) //set người dùng vào context
				}
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
