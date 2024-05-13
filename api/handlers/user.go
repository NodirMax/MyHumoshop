package handlers

import (
	"HumoSHOP/api/middleware"
	"HumoSHOP/api/response"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

// Обработчик отвечающий за Вход в систему (Авторизация)
func AuthorizationUzer(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 400,
		})
		return
	}

	token, err := services.AuthorizationUser(user)
	if err != nil {
        // Проверяем ошибку и устанавливаем соответствующий заголовок
        switch err.Error() {

		case "такого пользователя нет":
            response.ErrorJsonMessage(w, response.Resp{
				Message: "такого пользователя нет",
				StatusCode: 400,
			})
		
		case "ошибка в пароле или имени пользователя":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка в пароле или имени пользователя",
				StatusCode: 400,
			})
		
        case "ошибка на стороне сервера":
            response.ErrorJsonMessage(w, response.Resp{
				Message: "такого пользователя нет",
				StatusCode: 500,
			})

		case "ошибка при создании токена":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка при создании токена",
				StatusCode: 500,
			})
		
		case "ошибка при получении токена":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка при получении токена",
				StatusCode: 500,
			})
		}

        return
    }


	response.SuccessJsonMessage(w, response.Resp{
		Resp: token,
		Message: "Пользователь зашёл на сайт",
		StatusCode: http.StatusOK,
	})
	return
}	

// Обработчик отвечающий за Регистрацию
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
		})
		return
	}
	
	token, err := services.RegisterUser(user)
	if err != nil {
		// Проверяем ошибку и устанавливаем соответствующий заголовок
		switch err.Error() {
		case "поля имени не может быть пустым":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "поля имени не может быть пустым",
				StatusCode: 400,
			})

		case "пользователь с таким логином уже зарегистрирован":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "пользователь с таким логином уже зарегистрирован",
				StatusCode: 400,
			})

		case "длина логина должен быт не меньше 5 Символов":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "длина логина должен быт не меньше 5 Символов",
				StatusCode: 400,
			})

		case "длина пароля должен быт не меньше 5 Символов":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "длина пароля должен быт не меньше 5 Символов",
				StatusCode: 400,
			})
		
		case "ошибка при создание нового пользователя":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка при создание нового пользователя",
				StatusCode: 500,
			})
			
		case "ошибка при создании токена":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка при создании токена",
				StatusCode: 500,
			})
		}
		return
	}

	// Отправляем токен клиенту
	response.SuccessJsonMessage(w, response.Resp{
		Resp: token,
		Message: "Пользователь успешно зарегистрирован",
		StatusCode: http.StatusOK,
	})
	return
}

// Обработчик user->profile GET
func UserGET(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("login")

	user, err := services.GetUser(login)
	if err != nil{
		switch err.Error(){
		case "ошибка на стороне сервера":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка на стороне сервера",
				StatusCode: 500,
			})
			return
		}
	}
	
	response.SuccessJsonMessage(w, response.Resp{
		Resp: user,
		Message: "Данные пользователя успешно получени",
		StatusCode: 200,
	})
	return
}

//Обработчик user->profile PUT 
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "ошибка при получении данных",
			StatusCode: 400,
		})
	}
	
	user.Login = r.Header.Get("login")

    // запрос в пакет Servise
	err = services.UserUpdate(user)
	if err != nil{
		switch err.Error(){
		case "поля имени не может быть пустым":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "поля имени не может быть пустым",
				StatusCode: 400,
			})

		case "длина пароля должен быт не меньше 5 Символов":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "длина пароля должен быт не меньше 5 Символов",
				StatusCode: 400,
			})

		case "ошибка на стороне сервера":
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка на стороне сервера",
				StatusCode: 400,
			})
		}
		return
	}

	// Успешное выполнения запроса
	response.SuccessJsonMessage(w, response.Resp{
		Resp: user,
		Message: "Данные успешно обновлени",
		StatusCode: 200,
	})
	return
}

func UserGETAll(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w ,response.Resp{
			Message: "у вас нету достаточно прав",
			StatusCode: 403,
		})
		return
	}

	res, err := services.GetUserALL()
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 500,
		})
	}

	response.SuccessJsonMessage(w, response.Resp{
		Resp: res,
		Message: "Данные пользователей",
		StatusCode: 200,
	})
	return
}
// Обработчик user->profile/DELETE
// Нет