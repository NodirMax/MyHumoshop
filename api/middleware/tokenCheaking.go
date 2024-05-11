package middleware

import (
	"HumoSHOP/api/response"
	"HumoSHOP/pkg/utils"
	"net/http"
)
func ProtectedEndpoint(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Получаем токен из заголовка Authorization
	tokenString := r.Header.Get("Authorization")
	
	login, err := utils.ParseToken(tokenString)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Пользлватель не авторизован",
			StatusCode: 401,
		})
		return
	}


	r.Header.Set("login", login)
	h.ServeHTTP(w, r)
	})
}