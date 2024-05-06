package models

type ProductModel struct {
	Product_id    int64   `json:"product_id"`
	Product_name  string  `json:"product_name"`
	Product_price float64 `json:"product_price"`
	In_stock      bool    `json:"in_stock"`
	Category_id   int64   `json:"category_id"`
}