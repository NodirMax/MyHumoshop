package handlers

import (
	"HumoSHOP/api/middleware"
	"HumoSHOP/api/response"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CategoryGET(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryModel
	resp, err := services.CategoryGET(category)
	if err != nil {
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Ошибка при получении категории",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	// Получение результате от service
	response.SuccessJsonMessage(w, response.Resp{
		Resp:       resp,
		Message:    "Категория успешно получена",
		StatusCode: http.StatusOK,
	})
	return
}
 
// Получение данных о категории по id
func CategoryGETbyid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Некорректный идентификатор категории",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	// Получение данных от пакета service
	res, err := services.CategoryGETbyid(categoryID)
	if err != nil {
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Ошибка при получении категории",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	
	// Передаем обработчику результат
	response.SuccessJsonMessage(w, response.Resp{
		Resp:       res,
		Message:    "Категория успешно получена",
		StatusCode: http.StatusOK,
	})
}

// Создание новой категории 
func CategoryCreate(w http.ResponseWriter, r *http.Request)  {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Доступ запрещён!",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	var category models.CategoryModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Некоректный формат данных",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	err = services.CategoryCreate(category)
	if err != nil{
		if err.Error() == "поля имени категории не может быть пустым"{
			response.ErrorJsonMessage(w, response.Resp{
				Message:    "поля имени категории не может быть пустым",
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		if err.Error() == "поля категории с таким именем уже существует"{
			response.ErrorJsonMessage(w, response.Resp{
				Message:    "поля категории с таким именем уже существует",
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при создании категории",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	// Json Response
	response.SuccessJsonMessage(w, response.Resp{
		Message: "Категория успешно было добавлено",
		StatusCode: http.StatusOK,
	})
	return
}

// Обновления данных о категорый
func CategoryUpdate(w http.ResponseWriter, r *http.Request)  {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Доступ запрещён!",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	var category models.CategoryModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Некоректный формат данных",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	err = services.CategoryUpdate(category)
	if err != nil{
		if err.Error() == "поля имени категории не может быть пустым"{
			response.ErrorJsonMessage(w, response.Resp{
				Message:    "поля имени категории не может быть пустым",
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		if err.Error() == "поля категории с таким именем уже существует"{
			response.ErrorJsonMessage(w, response.Resp{
				Message:    "поля категории с таким именем уже существует",
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при обновлении категории",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.SuccessJsonMessage(w, response.Resp{
		Message: "Категория успешно было обновлени",
		StatusCode: http.StatusOK,
	})
	return
}

// Удаление данных о категории
func CategoryDELETE(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Доступ запрещён!",
			StatusCode: http.StatusForbidden,
		})
		return
	}

	var category models.CategoryModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Некоректный формат данных",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	
	err = services.CategoryDELETE(category)
	if err != nil{
		if err.Error() == "ошибка при передачи данных"{
			response.ErrorJsonMessage(w, response.Resp{
				Message:    "ошибка при передачи данных",
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при удалении категории",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	

	response.SuccessJsonMessage(w, response.Resp{
		Message: "Категория успешно удалена",
		StatusCode: http.StatusOK,
	})
	return
}
