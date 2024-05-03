package middleware

import (
	"HumoSHOP/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Задаем секретный ключ, который будет использоваться для подписи токена
var JwtKey = []byte("HumoShop")

// Структура для представления информации о пользователе, которая будет включена в токен
type Claims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

// Функция для создания JWT токена
func CreateToken(login string) (string, error) {
	// Устанавливаем срок действия токена на 1 час
	expirationTime := time.Now().Add(time.Hour)

	// Создаем структуру Claims
	claims := &Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Создаем токен с указанием алгоритма подписи и Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа и получаем строковое представление токена
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("ошибка создания токена: %v", err)
	}

	return tokenString, nil
}

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

func GenerateToken(login string) (string, error) {
	user, err := repository.CheckingUser(login)
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

func ExtractLoginFromToken(tokenString string) (string, error) {
    // Парсим токен
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })

    // Проверяем наличие ошибок
    if err != nil {
        return "", fmt.Errorf("ошибка при парсинге токена: %v", err)
    }

    // Проверяем валидность токена
    if !token.Valid {
        return "", fmt.Errorf("недействительный токен")
    }

    // Извлекаем информацию о логине из токена
    claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.Login, nil
	
}

func ParseToken(accessToken string) (string, error) {
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