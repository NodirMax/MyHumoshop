package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

// Получение данных о продукте
func ProductGetService(p models.ProductModel) (res models.ProductModel, err error) {
	res, err = repository.ProductGetDB(p.ProductID)
	if err != nil{
		return models.ProductModel{}, errors.New("ошибка получение данных о продукте из БД")
	}
	return 
}

// Добавление нового продукта
func ProductCreate(product models.ProductModel) (err error) {
	if product.ProductName == ""{
		return errors.New("поля name не может быть пустым")
	}

	if product.ProductPrice == 0{
		return errors.New("поля price не может быть пустым")
	}

	if product.CategoryID == 0{
		return errors.New("поля Category_id не может быть пустым")
	}


	err = repository.ProductCreateDB(product)
	if err != nil{
		return errors.New("ошибка со сторони сервера")
	}

	return
}

// Обновление данных о продукте
func ProductUpdate(product models.ProductModel) (err error) {
	err = repository.ProductUpdateDB(product)
	return
}

// Удаление данных продукта
func ProductDELETE(product models.ProductModel) (err error){
	err = repository.ProductDELETEDB(product.ProductID)
	return
}