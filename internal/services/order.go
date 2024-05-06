package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

func OrderPOSTService(orders []models.OrderModel) (summa float64, err error) {
	summa = 0
	for _, order := range orders{
		res, err := repository.ProductGetDB(order.Product_id)
		if err != nil{
			return 0, errors.New("ошибка при получение даннх")
		}
		summa += res.Product_price * float64(order.Product_count)
	}
	return
}