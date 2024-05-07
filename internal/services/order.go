package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

func OrderCreate(order models.OrderModel) (summa float64, err error) {
	summa = 0.0
	
	for _, product := range order.Products {
		price, err := repository.ProductGetDB(product.Product_id)
		if err != nil{
			return 0, errors.New("продукт не найден")
		}
		summa += float64(product.Product_count) * price.Product_price
	}
	return
}