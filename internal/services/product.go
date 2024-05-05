package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

// Получение данных о продукте
func ProductGetService(p models.ProductModels) (product models.ProductModels, err error) {
	res, err := repository.ProductGetDB(p.Product_id)
	if err != nil{
		return models.ProductModels{}, errors.New("ошибка получение данных о продукте из БД")
	}
	return res, nil
}

// Добавление нового продукта
func ProductPOSTService(product models.ProductModels) (err error) {
	if product.Product_name == ""{
		return errors.New("поля name не может быть пустым")
	}

	if product.Product_price == 0{
		return errors.New("поля price не может быть пустым")
	}

	if product.Category_id == 0{
		return errors.New("поля Category_id не может быть пустым")
	}


	err = repository.ProductPOSTDB(product)
	if err != nil{
		return errors.New("ошибка со сторони сервера")
	}

	return
}

// Обновление данных о продукте
func ProductPUTService(product models.ProductModels) (err error) {
	err = repository.ProductPUTDB(product)
	return
}

// Удаление данных продукта
func ProductDELETEService(product models.ProductModels) (err error){
	err = repository.ProductDELETEDB(product.Product_id)
	return
}