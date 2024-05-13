package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"database/sql"
	"errors"
	"log"
)

// Получение данных о продукте
func ProductGetDB(ProductID int64) (product models.ProductModel, err error) {
	row := db.DB.QueryRow("SELECT * FROM product WHERE product_id=$1", ProductID)
	err = row.Scan(&product.ProductID, &product.ProductName, &product.ProductPrice, &product.InStock, &product.CategoryID, &product.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
            return models.ProductModel{}, nil
        }
		return models.ProductModel{}, err
	}
	return 
}


// Получение данных о бо всех продуктах
func ProductGetALLDB() (products []models.ProductModel, err error) {
	rows, err := db.DB.Query("SELECT * FROM product")
	if err != nil{
		return []models.ProductModel{}, errors.New("ошибка при получении данных")
	}
	defer rows.Close()
	for rows.Next(){
		var p models.ProductModel
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.InStock, &p.CategoryID, &p.CategoryName)
		if err != nil{
			log.Println(err)
			continue
		}
		products = append(products, p)
	}
	return
}


// Добавления нового продукта
func ProductCreateDB(product models.ProductModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO product(product_name, product_price, in_stock, category_id, category_name) VALUES($1, $2, $3, $4, $5)", product.ProductName, product.ProductPrice, product.InStock, product.CategoryID, product.CategoryName)
	return
}

// Изменение данных о продукте
func ProductUpdateDB(product models.ProductModel) (err error) {
	_, err = db.DB.Exec("UPDATE product SET product_name=$1, product_price=$2, in_stock=$3, category_id=$4, category_name=$5 WHERE product_id=$6", product.ProductName, product.ProductPrice, product.InStock, product.CategoryID, product.CategoryName, product.ProductID)
	return 
}

// Удаление данных о продукте
func ProductDELETEDB(productID int64) (err error) {
	_, err = db.DB.Exec("DELETE FROM product WHERE product_id=$1", productID) 
	return
}