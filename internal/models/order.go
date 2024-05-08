package models

type OrderModel struct {
	Order_id int64 `json:"order_id"`
	User_id  int64 `json:"user_id"`
	Products []struct {
		Product_name  string `json:"product_name"`
		Product_id    int64  `json:"product_id"`
		Product_count int64  `json:"product_count"`
	} `json:"products"`
	Datatime string `json:"Datatime"`
}