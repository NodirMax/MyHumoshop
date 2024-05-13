package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

func OrderCreate(order models.OrderModel) (err error) {
	summa := 0.0
	if order.UserID != 0{
		return errors.New("ошибка при получении данных")
	}
	
	for _, product := range order.Products {
		price, err := repository.ProductGetDB(product.ProductID)
		if err != nil{
			return errors.New("продукт не найден")
		}
		summa += float64(product.ProductCount) * price.ProductPrice
	}

	lastid, err := repository.OrderCreateDB(order, summa)
	if err != nil{
		return err
	}

	// добавление продуктов в таблицу 2
	for _, product := range order.Products{
		err = repository.OrderProductCreateDB(lastid, product.ProductID, product.ProductCount)
		if err != nil{
			return err
		}
	}
	return nil
}

func OrderGet(userID int64) (order []repository.OrderWithProducts, err error) {
	order, err = repository.OrderGetDB(userID)
	return
}

func OrderGetALL() (history []repository.OrderWithProducts, err error){
	history, err = repository.OrderGETALL()
	return
}