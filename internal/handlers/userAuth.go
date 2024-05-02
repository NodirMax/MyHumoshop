package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"errors"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserModule
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при декодировке данных"))
		return
	}
	
	err = services.RegisterUserService(user)
	if err != nil {
        // Проверяем ошибку и устанавливаем соответствующий заголовок
        switch err {

        case errors.New("пользователь с таким логином уже зарегистрирован"):
			w.Write([]byte("пользователь с таким логином уже зарегистрирован"))
            w.WriteHeader(http.StatusConflict) // 409 Conflict
		
        case errors.New("ошибка при создание нового пользователя"):
            w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		
        default:
            w.WriteHeader(http.StatusInternalServerError) // Если ошибка неизвестна, ставим 500
        }
        return
    }

	w.WriteHeader(200)
	w.Write([]byte("Успешняя регистрация"))
	return
}