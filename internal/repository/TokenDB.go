package repository

import "HumoSHOP/pkg/db"

// Добавление токена в Б.Д
func CreateTokenDB(user_id int64, token string) (err error) {
	_, err = db.DB.Query("INSERT INTO usertoken(user_id, token) VALUES($1, $2)", user_id, token)
	return
}
