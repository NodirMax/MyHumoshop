package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"log"
)

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

func Category_id_GETService(category_id int64) (product []models.ProductModels, err error) {
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