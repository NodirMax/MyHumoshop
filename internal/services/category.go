package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
)

func CategoryGETService(c models.CategoryModels) (category []models.CategoryModels, err error) {
	res, err := repository.CategoryGETDB()
	if err != nil{
		return ([]models.CategoryModels{}), err
	}
	
	return res, err
}

func Category_id_GETService(category models.CategoryModels) (product []models.ProductModels, err error) {
	res, err := repository.Category_id_GETService(category.Category_id)
	if err != nil{
		return ([]models.ProductModels{}), err
	}
	// log.Println(res)
	return res, nil
}