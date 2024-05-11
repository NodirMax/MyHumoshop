package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
	"unicode/utf8"
)

// Авторизация пользователя
func AuthorizationUser(u models.UserModel) (token string,err error){
	
	inf, err := repository.GetUserFromDB(u.Login)
	if err != nil{
		return "", errors.New("ошибка на стороне сервера")
	}

	// проверяем что мы получили данные пользователя с Б.Д.
	if inf == (models.UserModel{}){
		return "", errors.New("такого пользователя нет")
	}

	// Проверим пароль пользователя, после расхеширования
	z :=  utils.HeshChecking(inf.Password, u.Password)
	if z != nil{
		return "", errors.New("такого пользователя нет")
	}

	// Создаем новый токен
	token, err = utils.GenerateToken(u.Login)
	if err != nil{
		return "", errors.New("ошибка при создании токена")
	}

	return token, nil
}


// Регистрация пользователя
func RegisterUser(user models.UserModel) (token string,err error) {
	// Добавление нового пользоватля
	
	// Проверяем что поля name не пуст
	if user.Name == ""{
		return "", errors.New("поля имени не может быть пустым")
	}
	
	
	// Проверка наличия пользователья в Б.Д.
	inf, err := repository.GetUserFromDB(user.Login)
	if err != nil{
		return "", errors.New("ошибка со стороны сервера")
	}

	if inf != (models.UserModel{}) {
		return "", errors.New("пользователь с таким логином уже зарегистрирован")
	}
	
	// Проверяем Длину login
	if utf8.RuneCountInString(user.Login) < 5 {
		return "", errors.New("длина логина должен быт не меньше 5 Символов")
	}
	
	
	// Проверяем длину паролья
	if utf8.RuneCountInString(user.Password) < 5 {
		return "", errors.New("длина пароля должен быт не меньше 5 Символов")
	}


	user.Password = utils.Heshing(user.Password)
	err = repository.CreateNewUserToDB(user)
	if err != nil{
		return "", errors.New("ошибка при создание нового пользователя")
	}
	
	// Добавление нового токена
	token, err = utils.GenerateToken(user.Login)
	if err != nil{
		return "", errors.New("ошибка при создании токена")
	}

	return token, nil
}


func GetUser(login string) (user models.UserModel, err error) {
	user, err = repository.GetUserFromDB(login)
	if err != nil{
		return (models.UserModel{}), errors.New("ошибка на стороне сервера")
	}
	return 
}

func UserUpdate(user models.UserModel) (err error) {
	// Проверка имени пользователя
	if user.Name == ""{
		return errors.New("поля имени не может быть пустым")
	}

	// Проверка пароля пользователя
	if utf8.RuneCountInString(user.Password) < 5 {
		return errors.New("длина пароля должен быт не меньше 5 Символов")
	} 

	user.Password = utils.Heshing(user.Password)
	err = repository.UpdateUserToDB(user)
	if err != nil{
		return errors.New("ошибка на стороне сервера")
	}
	return
}

