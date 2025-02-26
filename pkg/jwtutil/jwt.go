package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var globalJWT *JWT

type JWTConfig struct {
	AccessTokenExpirationTime  time.Duration `koanf:"access_token_expiration_time_ns"`
	RefreshTokenExpirationTime time.Duration `koanf:"refresh_token_expiration_time_ns"`
	SecretKey                  string        `koanf:"secret_key"`
}

type JWT struct {
	config JWTConfig
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func Init(cfg JWTConfig) {
	jwt := &JWT{
		config: cfg,
	}

	globalJWT = jwt
}

func GenerateTokenPair(username string, email string) (TokenPair, error) {
	accessToken, err := generateToken(username, email, globalJWT.config.AccessTokenExpirationTime, AccessToken)
	if err != nil {
		return TokenPair{}, err
	}

	refreshToken, err := generateToken(username, email, globalJWT.config.RefreshTokenExpirationTime, RefreshToken)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func Generate(username string, email string) (string, error) {
	return generateToken(username, email, globalJWT.config.AccessTokenExpirationTime, AccessToken)
}

func generateToken(username string, email string, expiration time.Duration, tokenType TokenType) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(expiration).Unix(),
		"iat":      time.Now().Unix(),
		"type":     string(tokenType),
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
			return nil, errors.New("invalid token signing method")
		}
		return []byte(globalJWT.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func Verify(tokenString string, expectedType TokenType) (map[string]interface{}, error) {
	claims, err := Decode(tokenString)
	if err != nil {
		return nil, err
	}

	if tokenType, ok := claims["type"].(string); !ok || TokenType(tokenType) != expectedType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}

func RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := Decode(refreshToken)
	if err != nil {
		return "", err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid refresh token: missing username claim")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid refresh token: missing email claim")
	}

	return generateToken(username, email, globalJWT.config.AccessTokenExpirationTime, AccessToken)
}
