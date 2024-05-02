package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"log"
)

// Получение данных про пользователя из Б.Д.
func GetUserFromDB(login string) (user models.UserModule,err error) {
	row := db.DB.QueryRow(`SELECT * FROM users WHERE user_id=$1`, login)

	err = row.Scan(&user.Id, &user.Name, &user.Login, &user.Password)
	if err != nil {
		log.Println("Ошибка row")
	}
	return
}

// Создание нового пользователя в Базу данных
func CreateNewUserToDB(user models.UserModule) (err error) {
	_, err = db.DB.Exec("INSERT INTO users(name, login, password) VALUES($1, $2, $3)", user.Name, user.Login, user.Password)
	return
}