package models

type OrderModel struct {
	OrderID  int64 `json:"order_id"`
	UserID   int64 `json:"user_id"`
	Products []struct {
		ProductName  string `json:"product_name"`
		ProductID    int64  `json:"product_id"`
		ProductCount int64  `json:"product_count"`
	} `json:"products"`
	Datatime string `json:"Datatime"`
}