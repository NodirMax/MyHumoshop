package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

// Обработка множеств категорий
func CategoryGETService(c models.CategoryModels) (category []models.CategoryModels, err error) {
	res, err := repository.CategoryGETDB()
	if err != nil{
		return ([]models.CategoryModels{}), err
	}
	
	return res, err
}

// Обработка категории по id
func Category_id_GETService(category models.CategoryModels) (product []models.ProductModels, err error) {
	res, err := repository.Category_id_GETDB(category.Category_id)
	if err != nil{
		return ([]models.ProductModels{}), err
	}
	// log.Println(res)
	return res, nil
}

// Добавление новой категории
func Category_id_POSTService(category models.CategoryModels) (err error) {
	if category.Category_name == ""{
		return errors.New("поля имени категории не может быть пустым")
	}
	err = repository.Category_id_POSTDB(category)
	return
}

// Обновление данных о новой категории
func Category_id_PUTService(category models.CategoryModels) (err error) {
	if category.Category_name == ""{
		return errors.New("поля имени категории не может быть пустым")
	}
	err = repository.Category_id_PUTDB(category)
	if err != nil{
		return errors.New("ошибка обновления данных")
	}
	return
}

// Удаление данных о категории
func Category_id_DELETEService(category models.CategoryModels) (err error) {
	if category.Category_id == 0{
		return errors.New("ошибка при передачи данных")
	}
	err = repository.Category_id_DELETEDB(category.Category_id)
	if err != nil{
		return errors.New("ошибка при удалении категории")
	}
	return
}