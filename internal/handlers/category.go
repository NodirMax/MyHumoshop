package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

func CategoryGET(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryModels 
	resp, err := services.CategoryGETService(category)
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
func CategoryGET_id(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryModels

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на стороне сервера"))
		return
	}
	
	// Получение данных от пакета service
	res, err := services.Category_id_GETService(category)
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
func CategoryPOST_id(w http.ResponseWriter, r *http.Request)  {
	// проверяем админ ли пользователь 
	if r.Header.Get("login") != "admin"{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var category models.CategoryModels
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	err = services.Category_id_POSTService(category)
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
func CategoryPUT_id(w http.ResponseWriter, r *http.Request)  {
	// проверяем админ ли пользователь 
	if r.Header.Get("login") != "admin"{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var category models.CategoryModels
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	err = services.Category_id_PUTService(category)
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
func CategoryDELETE_id(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	if r.Header.Get("login") != "admin"{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var category models.CategoryModels
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}
	
	err = services.Category_id_DELETEService(category)
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