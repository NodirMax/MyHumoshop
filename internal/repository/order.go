package repository

import (
	"HumoSHOP/pkg/db"
	"log"
)

func OrderCreateDB(user_id int64, product []byte) (err error) {
	// log.Println(string(product))
	_, err = db.DB.Exec("INSERT INTO orders(user_id, products) VALUES ($1, $2)", user_id, product)
	log.Println(err)
	return
}