package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

func OrderProductCreateDB(orderID, productID, productCount int64) (err error){
	_, err = db.DB.Exec("INSERT INTO orderproducts(orderid, productid, product_count) VALUES ($1, $2, $3)", orderID, productID, productCount)
	if err != nil{
		return errors.New("ошибка при добавление данных в таблице orderproducts")
	}
	return
}

func OrderCreateDB(order models.OrderModel, totalcost float64) (lastID int64, err error) {
    row := db.DB.QueryRow("INSERT INTO orders (user_id, totalcost) VALUES ($1, $2) RETURNING order_id", order.UserID, totalcost)
	err = row.Scan(&lastID)
    if err != nil {
        return 0, errors.New("ошибка при получении id")
    }

    return lastID, nil
}

func OrderGetDB(userID int64) (order []models.OrderModel ,err error) {
	rows, err := db.DB.Query(`
	SELECT o.order_id, o.user_id, o.totalcost, o.datatime, op.productid, op.product_count
	FROM orders o
	JOIN orderproducts op ON o.order_id = op.orderid
   	WHERE o.user_id=$1
	`, userID)
	if err != nil{
		log.Println(err)
		return []models.OrderModel{}, err
	}
	defer rows.Close()
	log.Println(rows)
	// for rows.Next()

	// log.Println(rows)
	// var models
	// for rows.Next(){
		
	// }
	return []models.OrderModel{}, nil
}


