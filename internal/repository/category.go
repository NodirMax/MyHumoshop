package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"log"
)

// Получение данных о категориях из БД
func CategoryGETDB() (category []models.CategoryModels, err error) {
	rows, err := db.DB.Query("SELECT * FROM category")
	if err != nil {
		return
	}
	defer rows.Close()

	var c models.CategoryModels
	for rows.Next() {
		err := rows.Scan(&c.Category_id, &c.Category_name)
		if err != nil {
			log.Println("Ошибка row")
			continue
		}
		category = append(category, c)
	}
	return
}

// Получение данных о категории из БД
func Category_id_GETDB(category_id int64) (product []models.ProductModels, err error) {
	rows, err := db.DB.Query("SELECT * FROM product WHERE category_id=$1", category_id)
	if err != nil {
		return
	}
	defer rows.Close()

	var p models.ProductModels 
	for rows.Next() {
		err := rows.Scan(&p.Product_id, &p.Product_name, &p.Product_price, &p.In_stock, &p.Category_id)
		if err != nil {
			log.Println("Ошибка row")
			continue
		}
		product = append(product, p)
	}
	return
}

// Добавление новой категории в БД
func Category_id_POSTDB(category models.CategoryModels) (err error) {
	_, err = db.DB.Exec("INSERT INTO category(category_name) VALUES($1)", category.Category_name)
	return 
}

// Обновление данных о категории
func Category_id_PUTDB(category models.CategoryModels) (err error) {
	_, err = db.DB.Exec("UPDATE category SET category_name=$1 WHERE category_id=$2", category.Category_name, category.Category_id)
	return
}

// Удаление данных о категории
func Category_id_DELETEDB(category_id int64) (err error) {
	_, err = db.DB.Exec("DELETE FROM category WHERE category_id=$1", category_id)
	return
}