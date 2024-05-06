package models

type OrderModel struct {
	Order_id      int64  `json:"order_id"`
	User_id       int64  `json:"user_id"`
	Product_id    int64  `json:"product_id"`
	Product_count int64  `json:"product_count"`
	Datatime      string `json:"Datatime"`
}