package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"log"
)

// Получение данных о категориях из БД
func CategoryGETDB() (category []models.CategoryModel, err error) {
	rows, err := db.DB.Query("SELECT * FROM category")
	if err != nil {
		return
	}
	defer rows.Close()

	var c models.CategoryModel
	for rows.Next() {
		err := rows.Scan(&c.CategoryID, &c.CategoryName)
		if err != nil {
			log.Println("Ошибка row")
			continue
		}
		category = append(category, c)
	}
	return
}

// Получение данных о категории из БД
func CategoryGETbyidDB(categoryID int) (product []models.ProductModel, err error) {
	rows, err := db.DB.Query("SELECT * FROM product WHERE category_id=$1", categoryID)
	if err != nil {
		return
	}
	defer rows.Close()

	var p models.ProductModel
	for rows.Next() {
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.InStock, &p.CategoryID, &p.CategoryName)
		if err != nil {
			log.Println("Ошибка row")
			continue
		}
		product = append(product, p)
	}
	return
}

// Получении категории по имени
func CategoryGETbyNameDB(CategoryName string) (ID int64) {
	_ = db.DB.QueryRow("SELECT category_id FROM category WHERE category_name=$1", CategoryName).Scan(&ID)
	return
}

// Добавление новой категории в БД
func CategoryCreateDB(category models.CategoryModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO category(category_name) VALUES($1)", category.CategoryName)
	return 
}

// Обновление данных о категории
func CategoryUpdateDB(category models.CategoryModel) (err error) {
	_, err = db.DB.Exec("UPDATE category SET category_name=$1 WHERE category_id=$2", category.CategoryName, category.CategoryID)
	return
}

// Удаление данных о категории
func CategoryDELETEDB(categoryID int64) (err error) {
	_, err = db.DB.Exec("DELETE FROM category WHERE category_id=$1", categoryID)
	return
}