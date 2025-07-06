package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenToken(email string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	expStr := os.Getenv("JWT_EXPIRATION")

	expSecs, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		log.Println("cannot parse JWT_EXPIRATION to number")
		return "", err
	}

	// JWT "exp" phải là số giây kể từ Unix epoch
	exp := time.Now().Unix() + expSecs

	claims := jwt.MapClaims{
		"jti": uuid.NewString(), // thêm ID duy nhất cho token
		"sub": email,
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra thuật toán ký
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	// Ép kiểu claims về MapClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
