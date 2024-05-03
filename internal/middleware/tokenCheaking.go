package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)
func ProtectedEndpoint(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Получаем токен из заголовка Authorization
	tokenString := r.Header.Get("Authorization")
	
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Требуется токен авторизации"))
		return
	}
	
	// Парсим и проверяем токен
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println(string(JwtKey))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка токена"))
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Если токен валидный, можно продолжить обработку запроса.
	// В данном примере мы просто возвращаем имя пользователя из токена.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h.ServeHTTP(w, r)
	})
}