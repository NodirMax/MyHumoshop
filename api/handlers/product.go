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

// Получение данных о продукте
func ProductGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response.ErrorJsonMessage(w, response.Resp{
			Message:    "Некорректный идентификатор категории",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	res, err := services.ProductGet(productID)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка со стороны сервера",
			StatusCode: 500,
		})
		return
	}
	
	response.SuccessJsonMessage(w, response.Resp{
		Resp: res,
		Message: "Данные о продукте",
		StatusCode: 200,
	})
	return
}


// Получение всех продуктов
func ProductGETALL(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w ,response.Resp{
			Message: "у вас недостаточно прав",
			StatusCode: 403,
		})
		return
	}
	
	products, err := services.ProductGETALL()
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 500,
		})
	}

	response.SuccessJsonMessage(w, response.Resp{
		Resp: products,
		Message: "Данные успешно получени",
		StatusCode: 200,
	})
	return
}


// Добавление нового продукта
func ProductCreate(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w ,response.Resp{
			Message: "у вас нету достаточно прав",
			StatusCode: 403,
		})
		return
	}

	var product models.ProductModel
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		response.ErrorJsonMessage(w ,response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 400,
		})
		return
	}

	err = services.ProductCreate(product)
	if err != nil{
		switch err.Error(){
		case "поля name не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля name не может быть пустым",
				StatusCode: 400,
			})
		
		case "поля price не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля price не может быть равним 0",
				StatusCode: 400,
			})
			
		case "поля Category_id не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля Category_id не может быть 0",
				StatusCode: 400,
			})
			
		case "поля Category_name не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля Category_name не может быть пустым",
				StatusCode: 400,
			})
			
		default:
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "Ошибка со стороны сервера",
				StatusCode: 500,
			})
			
		}
		return
	}
	
	response.SuccessJsonMessage(w, response.Resp{
		Resp: product,
		Message: "Продукт успешно добавлен",
		StatusCode: 200,
	})
}

// Обновление данных о продукте
func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w ,response.Resp{
			Message: "у вас нету достаточно прав",
			StatusCode: 403,
		})
		return
	}

	var product models.ProductModel
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 400,
		})
		return
	}

	err = services.ProductUpdate(product)
	if err != nil{
		switch err.Error(){
		case "поля name не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля name не может быть пустым",
				StatusCode: 400,
			})
		
		case "поля price не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля price не может быть равним 0",
				StatusCode: 400,
			})
			
		case "поля Category_id не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля Category_id не может быть 0",
				StatusCode: 400,
			})
			
		case "поля Category_name не может быть пустым":
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "поля Category_name не может быть пустым",
				StatusCode: 400,
			})
			
		default:
			response.ErrorJsonMessage(w ,response.Resp{
				Message: "Ошибка со стороны сервера",
				StatusCode: 500,
			})
			
		}
		return
	}

	response.SuccessJsonMessage(w, response.Resp{
		Resp: product,
		Message: "Данные о продукте успешно обновлени",
		StatusCode: 200,
	})
	return
}

// Удаление данных о продукте
func ProductDELETE(w http.ResponseWriter, r *http.Request) {
	// проверяем админ ли пользователь 
	login := r.Header.Get("login")
	if middleware.AdminCheck(login) != nil{
		response.ErrorJsonMessage(w ,response.Resp{
			Message: "у вас нету достаточно прав",
			StatusCode: 403,
		})
		return
	}

	var product models.ProductModel
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка при получении данных",
			StatusCode: 400,
		})
		return
	}

	err = services.ProductDELETE(product)
	if err != nil{
		response.ErrorJsonMessage(w, response.Resp{
			Message: "Ошибка на стороне сервера",
			StatusCode: 400,
		})
		return
	}

	response.SuccessJsonMessage(w, response.Resp{
		Message: "данные о продукте успешно удалены",
		StatusCode: 200,
	})
	return
}