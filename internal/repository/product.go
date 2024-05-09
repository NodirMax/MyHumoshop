package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"database/sql"
)

// Получение данных о продукте
func ProductGetDB(ProductID int64) (product models.ProductModel, err error) {
	row := db.DB.QueryRow("SELECT * FROM product WHERE product_id=$1", ProductID)
	err = row.Scan(&product.ProductID, &product.ProductName, &product.ProductPrice, &product.InStock, &product.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
            return models.ProductModel{}, nil
        }
		return models.ProductModel{}, err
	}
	return 
}
// Добавления нового продукта
func ProductCreateDB(product models.ProductModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO product(product_name, product_price, in_stock, category_id) VALUES($1, $2, $3, $4)", product.ProductName, product.ProductPrice, product.InStock, product.CategoryID)
	return
}

// Изменение данных о продукте
func ProductUpdateDB(product models.ProductModel) (err error) {
	_, err = db.DB.Exec("UPDATE product SET product_name=$1, product_price=$2, in_stock=$3, category_id=$4 WHERE product_id=$5", product.ProductName, product.ProductPrice, product.InStock, product.CategoryID, product.ProductID)
	return 
}

// Удаление данных о продукте
func ProductDELETEDB(productID int64) (err error) {
	_, err = db.DB.Exec("DELETE FROM product WHERE product_id=$1", productID) 
	return
}