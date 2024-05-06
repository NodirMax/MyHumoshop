package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"database/sql"
)

// Получение данных о продукте
func ProductGetDB(Product_id int64) (product models.ProductModel, err error) {
	row := db.DB.QueryRow("SELECT * FROM product WHERE product_id=$1", Product_id)
	err = row.Scan(&product.Product_id, &product.Product_name, &product.Product_price, &product.In_stock, &product.Category_id)
	if err != nil {
		if err == sql.ErrNoRows {
            return models.ProductModel{}, nil
        }
		return models.ProductModel{}, err
	}
	return 
}
// Добавления нового продукта
func ProductPOSTDB(product models.ProductModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO product(product_name, product_price, in_stock, category_id) VALUES($1, $2, $3, $4)", product.Product_name, product.Product_price, product.In_stock, product.Category_id)
	return
}

// Изменение данных о продукте
func ProductPUTDB(product models.ProductModel) (err error) {
	_, err = db.DB.Exec("UPDATE product SET product_name=$1, product_price=$2, in_stock=$3, category_id=$4 WHERE product_id=$5", product.Product_name, product.Product_price, product.In_stock, product.Category_id, product.Product_id)
	return 
}

// Удаление данных о продукте
func ProductDELETEDB(product_id int64) (err error) {
	_, err = db.DB.Exec("DELETE FROM product WHERE product_id=$1", product_id) 
	return
}