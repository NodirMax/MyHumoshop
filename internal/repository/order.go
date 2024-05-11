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

type OrderWithProducts struct {
    OrderID    int64
    UserID     int64
    TotalCost  float64
    Datetime   string
    Products   []Product
}

type Product struct {
    ProductID    int64
    ProductCount int64
}

func OrderGetDB(userID int64) ([]OrderWithProducts, error) {
    rows, err := db.DB.Query(`
        SELECT o.order_id, o.user_id, o.totalcost, o.datatime, op.productid, op.product_count
        FROM orders o
        JOIN orderproducts op ON o.order_id = op.orderid
        WHERE o.user_id=$1
    `, userID)
    
    if err != nil {
        log.Println(err)
        return nil, err
    }
    defer rows.Close()

    orders := make(map[int64]*OrderWithProducts)

    for rows.Next() {
        var orderID int64
        var userID int64
        var totalCost float64
        var datetime string
        var productID int64
        var productCount int64

        err := rows.Scan(&orderID, &userID, &totalCost, &datetime, &productID, &productCount)
        if err != nil {
            log.Println(err)
            continue
        }

        if order, ok := orders[orderID]; ok {
            order.Products = append(order.Products, Product{
                ProductID:    productID,
                ProductCount: productCount,
            })
        } else {
            orders[orderID] = &OrderWithProducts{
                OrderID:   orderID,
                UserID:    userID,
                TotalCost: totalCost,
                Datetime:  datetime,
                Products:  []Product{{ProductID: productID, ProductCount: productCount}},
            }
        }
    }

    result := make([]OrderWithProducts, 0, len(orders))
    for _, order := range orders {
        result = append(result, *order)
    }
    return result, nil
}

func OrderGETALL() ([]OrderWithProducts, error) {
    rows, err := db.DB.Query(`
        SELECT o.order_id, o.user_id, o.totalcost, o.datatime, op.productid, op.product_count
        FROM orders o
        JOIN orderproducts op ON o.order_id = op.orderid
    `)
    
    if err != nil {
        log.Println(err)
        return nil, err
    }
    defer rows.Close()

    orders := make(map[int64]*OrderWithProducts)

    for rows.Next() {
        var orderID int64
        var userID int64
        var totalCost float64
        var datetime string
        var productID int64
        var productCount int64

        err := rows.Scan(&orderID, &userID, &totalCost, &datetime, &productID, &productCount)
        if err != nil {
            log.Println(err)
            continue
        }

        if order, ok := orders[orderID]; ok {
            order.Products = append(order.Products, Product{
                ProductID:    productID,
                ProductCount: productCount,
            })
        } else {
            orders[orderID] = &OrderWithProducts{
                OrderID:   orderID,
                UserID:    userID,
                TotalCost: totalCost,
                Datetime:  datetime,
                Products:  []Product{{ProductID: productID, ProductCount: productCount}},
            }
        }
    }

    result := make([]OrderWithProducts, 0, len(orders))
    for _, order := range orders {
        result = append(result, *order)
    }
    return result, nil
}


