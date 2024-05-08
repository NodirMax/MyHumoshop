package handlers

import (
	"HumoSHOP/internal/middleware"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

func CategoryGET(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryModel
	resp, err := services.CategoryGET(category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на стороне сервера"))
		return
	}

	// Получение результате от service
	result, err := json.Marshal(resp)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write(result)
	return
}
 
// Получение данных о категории по id
func CategoryGETbyid(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryModel

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на стороне сервера"))
		return
	}
	
	// Получение данных от пакета service
	res, err := services.CategoryGETbyid(category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	rezult, err := json.Marshal(res)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}
	
	// Передаем обработчику результат
	w.WriteHeader(200)
	w.Write(rezult)
	return
}

// Создание новой категории 
func CategoryCreate(w http.ResponseWriter, r *http.Request)  {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheack(login) != nil{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var category models.CategoryModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	err = services.CategoryCreate(category)
	if err != nil{
		if err.Error() == "поля имени категории не может быть пустым"{
			w.WriteHeader(400)
			w.Write([]byte("поля имени категории не может быть пустым"))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Новая категория успешно добавлено"))
	return
}

// Обновления данных о категорый
func CategoryUpdate(w http.ResponseWriter, r *http.Request)  {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheack(login) != nil{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var category models.CategoryModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	err = services.CategoryUpdate(category)
	if err != nil{
		if err.Error() == "поля имени категории не может быть пустым"{
			w.WriteHeader(400)
			w.Write([]byte("поля имени категории не может быть пустым"))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("данные о категории успешно обновлени"))
	return
}

// Удаление данных о категории
func CategoryDELETE(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheack(login) != nil{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var category models.CategoryModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}
	
	err = services.CategoryDELETE(category)
	if err != nil{
		if err.Error() == "ошибка при передачи данных"{
			w.WriteHeader(400)
			w.Write([]byte("ошибка при передачи данных"))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные успешно удалени"))
	return
}