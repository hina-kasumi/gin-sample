package services

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenToken(email string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	expStr := os.Getenv("JWT_EXPIRATION")

	expSecs, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		log.Fatalln("cannot parse JWT_EXPIRATION to number")
		return "", err
	}

	// JWT "exp" phải là số giây kể từ Unix epoch
	exp := time.Now().Unix()

	claims := jwt.MapClaims{
		"email": email,
		"exp":   exp + expSecs,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
