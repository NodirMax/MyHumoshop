package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserModels
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при декодировке данных"))
		return
	}
	
	err = services.RegisterUserService(user)
	if err != nil {
        // Проверяем ошибку и устанавливаем соответствующий заголовок
        switch err.Error() {

		case "пользователь с таким логином уже зарегистрирован":
            w.WriteHeader(http.StatusConflict) // 409 Conflict
			w.Write([]byte("пользователь с таким логином уже зарегистрирован"))
		
        case "ошибка при создание нового пользователя":
            w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
			w.Write([]byte("ошибка на стороне сервера"))
		
		}

        return
    }

	w.WriteHeader(200)
	w.Write([]byte("Успешняя регистрация"))
	return
}