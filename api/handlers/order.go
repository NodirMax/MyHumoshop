package handlers

import (
	"HumoSHOP/api/response"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/services"
	"encoding/json"
	"net/http"
)

// Создаем список
func OrderCreate(w http.ResponseWriter, r *http.Request) {
	var order models.OrderModel 

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при полученнии данных",
			StatusCode: 400,
		})
	}

	err = services.OrderCreate(order)
	if err != nil{
		if err.Error() == "ошибка при получении данных"{
			response.ErrorJsonMessage(w, response.Resp{
				Message: "ошибка при получении данных",
				StatusCode: 400,
			})
			return
		}
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка на стороне сервера",
			StatusCode: 500,
		})
		return
	}

	response.ErrorJsonMessage(w, response.Resp{
		Message: "Данные успешно добавлени в БД",
		StatusCode: 200,
	})
	return
}

// Получаем список покупок пользователя
func OrderGet(w http.ResponseWriter, r *http.Request)  {
	login := r.Header.Get("login")
	user, err := services.GetUser(login)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при полученнии данных ",
			StatusCode: 400,
		})
		return
	}

	
	result, err := services.OrderGet(user.Id)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка на стороне сервера",
			StatusCode: 500,
		})
	}


	response.SuccessJsonMessage(w, response.Resp{
		Resp: result,
		Message: "История покупок",
		StatusCode: 200,
	})
	return
}

func OrderGetAll(w http.ResponseWriter, r *http.Request)  {
	res, err := services.OrderGetALL()
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 500,
		})
	}
	response.SuccessJsonMessage(w, response.Resp{
		Resp: res,
		Message: "Данные успешно получени",
		StatusCode: 200,
	})
}