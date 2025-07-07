package services

import (
	"errors"
	dtos "goprj/DTOs"
	"goprj/entities"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func LoginService(req dtos.LoginRequest) (string, error) {
	user := &entities.User{Email: req.Email, Password: req.Password}
	dbUser, err := FindOneUser(*user)

	if err != nil || dbUser.IsValidPassword(user.Password) != nil {
		return "", errors.New("invalid email or password")
	}

	tokenString, err := GenToken(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	return tokenString, err
}

func LogoutService(token string) error {
	return AddTokenToBlackList(token)
}

func AddTokenToBlackList(token string) error {
	claims, err := ParseAccessToken(token)

	if err != nil {
		log.Println("can not parse token")
		return err
	}

	jti := claims["jti"].(string)
	sub := claims["sub"].(string)
	key := os.Getenv("JWT_BLACKLIST_PREFIX") + jti + sub
	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid exp format")
	}
	exp := int64(expFloat)

	SetRedisExpire(key, "blacklist", exp)
	return nil
}

func IsTokenInBlackList(token string) bool {
	claims, err := ParseAccessToken(token)

	if err != nil {
		log.Println("can not parse token")
		return true
	}

	jti := claims["jti"].(string)
	sub := claims["sub"].(string)
	key := os.Getenv("JWT_BLACKLIST_PREFIX") + jti + sub

	val, err := GetRedisValue(key)

	if val == "blacklist" || (err != nil && err != redis.Nil) {
		log.Println("token in blacklist")
		return true
	}

	return false
}
