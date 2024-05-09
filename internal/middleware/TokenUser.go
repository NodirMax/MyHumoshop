package middleware

import (
	"HumoSHOP/config"
	"HumoSHOP/internal/repository"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Структура для представления информации о пользователе, которая будет включена в токен
type Claims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

const (
	tokenTTL   = 12 * time.Hour
)

func GenerateToken(login string) (string, error) {
	// берем секретный ключ из конфига
	signingKey := config.Settings.JWTSecret
	
	user, err := repository.GetUserFromDB(login)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Login,
	})

	return token.SignedString([]byte(signingKey))
}


func ParseToken(accessToken string) (string, error) {
	// берем секретный ключ из конфига
	signingKey := config.Settings.JWTSecret

	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Login, nil
}