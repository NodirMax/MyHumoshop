package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

// Обработка множеств категорий
func CategoryGET(c models.CategoryModel) (category []models.CategoryModel, err error) {
	res, err := repository.CategoryGETDB()
	if err != nil{
		return ([]models.CategoryModel{}), err
	}
	
	return res, err
}

// Обработка категории по id
func CategoryGETbyid(category models.CategoryModel) (product []models.ProductModel, err error) {
	res, err := repository.CategoryGETbyidDB(category.Category_id)
	if err != nil{
		return ([]models.ProductModel{}), err
	}
	// log.Println(res)
	return res, nil
}

// Добавление новой категории
func CategoryCreate(category models.CategoryModel) (err error) {
	if category.Category_name == ""{
		return errors.New("поля имени категории не может быть пустым")
	}
	err = repository.CategoryCreateDB(category)
	return
}

// Обновление данных о новой категории
func CategoryUpdate(category models.CategoryModel) (err error) {
	if category.Category_name == ""{
		return errors.New("поля имени категории не может быть пустым")
	}
	err = repository.CategoryUpdateDB(category)
	if err != nil{
		return errors.New("ошибка обновления данных")
	}
	return
}

// Удаление данных о категории
func CategoryDELETE(category models.CategoryModel) (err error) {
	if category.Category_id == 0{
		return errors.New("ошибка при передачи данных")
	}
	err = repository.CategoryDELETEDB(category.Category_id)
	if err != nil{
		return errors.New("ошибка при удалении категории")
	}
	return
}