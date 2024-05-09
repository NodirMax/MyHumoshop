package middleware

import (
	"net/http"
)
func ProtectedEndpoint(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Получаем токен из заголовка Authorization
	tokenString := r.Header.Get("Authorization")
	
	login, err := ParseToken(tokenString)
	if err != nil{
		w.WriteHeader(401)
		w.Write([]byte("Пользователь не обнаружен!"))
		return
	}


	r.Header.Set("login", login)
	h.ServeHTTP(w, r)
	})
}