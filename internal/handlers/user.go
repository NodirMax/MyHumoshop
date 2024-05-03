package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

// Обработчик отвечающий за Вход в систему (Авторизация)
func AuthorizationUzer(w http.ResponseWriter, r *http.Request) {
	var user models.UserModels
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при дешифровке данных"))
		return
	}

	token, err := services.AuthorizationUserService(user)
	if err != nil {
        // Проверяем ошибку и устанавливаем соответствующий заголовок
        switch err.Error() {

		case "такого пользователя нет":
            w.WriteHeader(http.StatusConflict) // 409 Conflict
			w.Write([]byte("Неверный пароль или имя пользователя"))
		
        case "ошибка на стороне сервера":
            w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
			w.Write([]byte("ошибка на стороне сервера"))

		case "ошибка при создании токена":
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ошибка при создании токена"))
		
		case "ошибка":
			w.WriteHeader(401)
			w.Write([]byte("Ошибка при получение токена"))
		}

        return
    }
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := token
	json.NewEncoder(w).Encode(response)
	return
}	

// Обработчик отвечающий за Регистрацию
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserModels
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при декодировке данных"))
		return
	}
	
	token, err := services.RegisterUserService(user)
	if err != nil {
		// Проверяем ошибку и устанавливаем соответствующий заголовок
		switch err.Error() {

		case "пользователь с таким логином уже зарегистрирован":
			w.WriteHeader(http.StatusConflict) // 409 Conflict
			w.Write([]byte("пользователь с таким логином уже зарегистрирован"))
		
		case "ошибка при создание нового пользователя":
			w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
			w.Write([]byte("ошибка на стороне сервера"))
		case "ошибка при создании токена":
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ошибка при создании токена"))
		}

		return
	}
	// Отправляем токен клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := token
	json.NewEncoder(w).Encode(response)

	w.Write([]byte("Успешняя регистрация"))
	return
}

// Обработчик user->profile
func UserGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SECRET PAGE"))
	return
}


