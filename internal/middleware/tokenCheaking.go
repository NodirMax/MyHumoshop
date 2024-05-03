package middleware

import (
	"net/http"
)
func ProtectedEndpoint(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Получаем токен из заголовка Authorization
	tokenString := r.Header.Get("Authorization")
	
	_, err := ParseToken(tokenString)
	if err != nil{
		w.WriteHeader(401)
		w.Write([]byte("Пользователь не обнаружен!"))
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h.ServeHTTP(w, r)
	})
}