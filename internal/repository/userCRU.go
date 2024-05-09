package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"database/sql"
)

// Получение данных про пользователя из Б.Д.
func GetUserFromDB(login string) (user models.UserModel, err error) {
	row := db.DB.QueryRow(`SELECT * FROM users WHERE login=$1`, login)
	err = row.Scan(&user.Id, &user.Name, &user.Login, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
            return user, nil
        }
		return 
	}
	return user, nil
}


// Создание нового пользователя в Базу данных
func CreateNewUserToDB(user models.UserModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO users(name, login, password) VALUES($1, $2, $3)", user.Name, user.Login, user.Password)
	return
}

func UpdateUserToDB(user models.UserModel) (err error) {
	_, err = db.DB.Exec("UPDATE users SET name=$1, password=$2 WHERE login=$3", user.Name, user.Password, user.Login)
	return 
}