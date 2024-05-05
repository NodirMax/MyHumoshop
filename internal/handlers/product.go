package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

// Получение данных о продукте
func ProductGet(w http.ResponseWriter, r *http.Request) {
	var product models.ProductModels
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}
	
	res, err := services.ProductGetService(product)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на стороне сервера"))
		return
	}
	
	result, err := json.Marshal(res)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}
	
	w.WriteHeader(200)
	w.Write(result)
	return
}


// Добавление нового продукта
func ProductPOST(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	if r.Header.Get("login") != "admin"{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var product models.ProductModels
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}

	err = services.ProductPOSTService(product)
	if err != nil{
		switch err.Error(){
		case "поля name не может быть пустым":
			w.WriteHeader(400)
			w.Write([]byte("поля name не может быть пустым"))
			return
		case "поля price не может быть пустым":
			w.WriteHeader(400)
			w.Write([]byte("поля price не может быть пустым"))
			return
		case "поля Category_id не может быть пустым":
			w.WriteHeader(400)
			w.Write([]byte("поля Category_id не может быть пустым"))
			return
		default:
			w.WriteHeader(500)
			w.Write([]byte("ошибка со стороны сервера"))
			return
		}
	}

	w.WriteHeader(200)
	w.Write([]byte("Новый продукт успешно добавлен"))
	return	
}

// Обновление данных о продукте
func ProductPUT(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	if r.Header.Get("login") != "admin"{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var product models.ProductModels
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}

	err = services.ProductPUTService(product)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные о продукте успешно обновлени"))
	return
}

// Удаление данных о продукте
func ProductDELETE(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	if r.Header.Get("login") != "admin"{
		w.WriteHeader(403)
		w.Write([]byte("вы не имеете достаточно прав"))
		return
	}

	var product models.ProductModels
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}

	err = services.ProductDELETEService(product)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("данные о продукте успешно удалены"))
	return
}