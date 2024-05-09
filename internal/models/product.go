package models

type ProductModel struct {
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	InStock      bool    `json:"in_stock"`
	CategoryID   int64   `json:"category_id"`
}