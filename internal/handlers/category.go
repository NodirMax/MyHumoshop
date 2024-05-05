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

	w.WriteHeader(200)
	w.Write(rezult)
	return
}