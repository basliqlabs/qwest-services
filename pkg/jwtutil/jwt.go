package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var globalJWT *JWT

type JWTConfig struct {
	AccessTokenExpirationTime time.Duration `koanf:"access_token_expiration_time_ns"`
	SecretKey                 string        `koanf:"secret_key"`
}

type JWT struct {
	config JWTConfig
}

func Init(cfg JWTConfig) {
	jwt := &JWT{
		config: cfg,
	}

	globalJWT = jwt
}

func Generate(username string, email string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(globalJWT.config.AccessTokenExpirationTime).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(globalJWT.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Decode(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("")
		}
		return []byte(globalJWT.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("")
	}
}
