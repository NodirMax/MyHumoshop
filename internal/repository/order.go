package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"errors"

	_ "github.com/lib/pq"
)

func OrderProductCreateDB(order_id, product_id, product_count int64) (err error){
	_, err = db.DB.Exec("INSERT INTO orderproducts(orderid, productid, product_count) VALUES ($1, $2, $3)", order_id, product_id, product_count)
	if err != nil{
		return errors.New("ошибка при добавление данных в таблице ordersproduct")
	}
	return
}

func OrderCreateDB(order models.OrderModel, totalcost float64) (lastID int64, err error) {
    row := db.DB.QueryRow("INSERT INTO orders (user_id, totalcost) VALUES ($1, $2) RETURNING order_id", order.User_id, totalcost)
	err = row.Scan(&lastID)
    if err != nil {
        return 0, errors.New("ошибка при получении id")
    }

    return lastID, nil
}



