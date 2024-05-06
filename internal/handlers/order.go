package handlers

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func OrderPOST(w http.ResponseWriter, r *http.Request) {
	var order []models.OrderModel 

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при декодирование"))
		return
	}

	summa, err := services.OrderPOSTService(order)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("ошибка на стороне сервера"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(strconv.FormatFloat(summa, 'f', -1, 64)))
	return
}