package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

// Обработка множеств категорий
func CategoryGET(c models.CategoryModel) (res []models.CategoryModel, err error) {
	res, err = repository.CategoryGETDB()
	if err != nil{
		return ([]models.CategoryModel{}), err
	}
	
	return 
}

// Обработка категории по id
func CategoryGETbyid(categoryID int) (res []models.ProductModel, err error) {
	res, err = repository.CategoryGETbyidDB(categoryID)
	if err != nil{
		return ([]models.ProductModel{}), err
	}
	if len(res) == 0 {
		res = []models.ProductModel{}
	}

	return 
}

// Добавление новой категории
func CategoryCreate(category models.CategoryModel) (err error) {
	// Проверяем что имени категории явлвяестя пустим или нет
	if category.CategoryName == ""{
		return errors.New("поля имени категории не может быть пустым")
	}

	// Проверим существует ли данная категория в БД уже
	id := repository.CategoryGETbyNameDB(category.CategoryName)
	if id != 0 {
		return errors.New("поля категории с таким именем уже существует")
	}

	err = repository.CategoryCreateDB(category)
	return
}

// Обновление данных о новой категории
func CategoryUpdate(category models.CategoryModel) (err error) {
	// Проверяем что имени категории явлвяестя пустим или нет
	if category.CategoryName == ""{
		return errors.New("поля имени категории не может быть пустым")
	}

	// Проверим существует ли данная категория в БД уже
	id := repository.CategoryGETbyNameDB(category.CategoryName)
	if id != 0 {
		return errors.New("поля категории с таким именем уже существует")
	}

	err = repository.CategoryUpdateDB(category)
	if err != nil{
		return errors.New("ошибка обновления данных")
	}
	return
}

// Удаление данных о категории
func CategoryDELETE(category models.CategoryModel) (err error) {
	// Категория ID не может быть равным 0
	if category.CategoryID == 0{
		return errors.New("ошибка при передачи данных")
	}

	err = repository.CategoryDELETEDB(category.CategoryID)
	if err != nil{
		return errors.New("ошибка при удалении категории")
	}
	
	return
}