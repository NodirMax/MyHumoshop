package repository

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/pkg/db"
	"database/sql"
	"log"
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
	return 
}


// Создание нового пользователя в Базe данных
func CreateNewUserToDB(user models.UserModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO users(name, login, password) VALUES($1, $2, $3)", user.Name, user.Login, user.Password)
	return
}

func UpdateUserToDB(user models.UserModel) (err error) {
	_, err = db.DB.Exec("UPDATE users SET name=$1, password=$2 WHERE login=$3", user.Name, user.Password, user.Login)
	return 
}

func GetUserALLDB() (users []models.UserModel, err error){
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil{
		return ([]models.UserModel{}), err
	}
	defer rows.Close()

	for rows.Next(){
	var u models.UserModel
	err := rows.Scan(&u.Id, &u.Name, &u.Login, &u.Password, &u.Role)
        if err != nil{
            log.Println(err)
            continue
        }
		users = append(users, u)
	}
	return
}