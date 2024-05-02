package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Задаем секретный ключ, который будет использоваться для подписи токена
var jwtKey = []byte("HumoShop")

// Структура для представления информации о пользователе, которая будет включена в токен
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func main() {
    // Пример создания токена
    token := CreateToken("example_user")
    fmt.Println("Token:", token)

    // Пример проверки токена
    username, err := VerifyToken(token)
    if err != nil {
        fmt.Println("Ошибка проверки токена:", err)
    } else {
        fmt.Println("Имя пользователя из токена:", username)
    }
}

// Функция для создания JWT токена
func CreateToken(username string) string {
    // Устанавливаем срок действия токена на 1 час
    expirationTime := time.Now().Add(time.Hour)

    // Создаем структуру Claims
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Создаем токен с указанием алгоритма подписи и Claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Подписываем токен с использованием секретного ключа и получаем строковое представление токена
    tokenString, err := token.SignedString(jwtKey)
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
        return jwtKey, nil
    })

    if err != nil {
        return "", err
    }

    // Проверяем валидность токена
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims.Username, nil
    } else {
        return "", fmt.Errorf("неверный токен")
    }
}
