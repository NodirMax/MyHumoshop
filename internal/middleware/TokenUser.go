package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Задаем секретный ключ, который будет использоваться для подписи токена
var JwtKey = []byte("HumoShop")

// Структура для представления информации о пользователе, которая будет включена в токен
type Claims struct {
    Login string `json:"login"`
    jwt.StandardClaims
}

// Функция для создания JWT токена
func CreateToken(login string) string {
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
        fmt.Println("Ошибка создания токена:", err)
        return ""
    }

    return tokenString
}

// Функция для проверки и расшифровки JWT токена
func VerifyToken(tokenString string) (string, error) {
    // Парсим токен
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })

    if err != nil {
        return "", err
    }

    // Проверяем валидность токена
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims.Login, nil
    } else {
        return "", fmt.Errorf("неверный токен")
    }
}
