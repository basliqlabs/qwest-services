package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

func New(cfg JWTConfig) JWT {
	return JWT{
		config: cfg,
	}
}

func (j *JWT) GenerateTokenPair(username string, email string) (TokenPair, error) {
	accessToken, err := j.generateToken(username, email, j.config.AccessTokenExpirationTime, AccessToken)
	if err != nil {
		return TokenPair{}, err
	}

	refreshToken, err := j.generateToken(username, email, j.config.RefreshTokenExpirationTime, RefreshToken)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *JWT) Generate(username string, email string) (string, error) {
	return j.generateToken(username, email, j.config.AccessTokenExpirationTime, AccessToken)
}

func (j *JWT) generateToken(username string, email string, expiration time.Duration, tokenType TokenType) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(expiration).Unix(),
		"iat":      time.Now().Unix(),
		"type":     string(tokenType),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) Decode(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(j.config.SecretKey), nil
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

func (j *JWT) Verify(tokenString string, expectedType TokenType) (map[string]interface{}, error) {
	claims, err := j.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	if tokenType, ok := claims["type"].(string); !ok || TokenType(tokenType) != expectedType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}

func (j *JWT) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := j.Verify(refreshToken, RefreshToken)
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

	return j.generateToken(username, email, j.config.AccessTokenExpirationTime, AccessToken)
}
