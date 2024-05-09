package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

// Создаем список
func OrderCreate(w http.ResponseWriter, r *http.Request) {
	var order models.OrderModel 

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}

	err = services.OrderCreate(order)
	if err != nil{
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные успешно добавлени в БД"))
	return
}

// Получаем список покупок пользователя
func OrderGet(w http.ResponseWriter, r *http.Request)  {
	var order models.OrderModel
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}
	_, err = services.OrderGet(order.UserID)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка со стороны сервера"))
		return
	}
	w.WriteHeader(200)
	// w.Write(order)
	return
}