package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func AuthorizationUzer(w http.ResponseWriter, r *http.Request) {
	var user models.UserModels
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при дешифровке данных"))
		return
	}

	err = services.AuthorizationUserService(user)
	if err != nil {
        // Проверяем ошибку и устанавливаем соответствующий заголовок
        switch err.Error() {

		case "такого пользователя нет":
            w.WriteHeader(http.StatusConflict) // 409 Conflict
			w.Write([]byte("Неверный пароль или имя пользователя"))
		
        case "ошибка на стороне сервера":
            w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
			w.Write([]byte("ошибка на стороне сервера"))
		}

        return
    }
	
	w.WriteHeader(200)
	fmt.Fprintf(w, "Пользователь %s вошёл в систему", user.Login)
	return
}