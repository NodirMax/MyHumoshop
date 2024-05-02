package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"database/sql"
)

// Получение данных про пользователя из Б.Д.
func GetUserFromDB(login string) (user models.UserModule,err error) {
	row := db.DB.QueryRow(`SELECT * FROM users WHERE login=$1`, login)

	err = row.Scan(&user.Id, &user.Name, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
            return user, nil
        }
		return 
	}
	return user, nil
}

// Создание нового пользователя в Базу данных
func CreateNewUserToDB(user models.UserModule) (err error) {
	_, err = db.DB.Exec("INSERT INTO users(name, login, password) VALUES($1, $2, $3)", user.Name, user.Login, user.Password)
	return
}