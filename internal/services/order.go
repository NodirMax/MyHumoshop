package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

func OrderCreate(order models.OrderModel) (err error) {
	summa := 0.0
	
	for _, product := range order.Products {
		price, err := repository.ProductGetDB(product.Product_id)
		if err != nil{
			return errors.New("продукт не найден")
		}
		summa += float64(product.Product_count) * price.Product_price
	}

	lastid, err := repository.OrderCreateDB(order, summa)
	if err != nil{
		return err
	}

	// добавление продуктов в таблицу 2
	for _, product := range order.Products{
		err = repository.OrderProductCreateDB(lastid, product.Product_id, product.Product_count)
		if err != nil{
			return err
		}
	}
	return nil
}