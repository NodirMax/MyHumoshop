package services

import (
	"HumoSHOP/internal/middleware"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

func GetUserFromService(login string) ( user models.UserModel,err error) {
	user, err = repository.GetUserFromDB(login)
	if err != nil{
		return (models.UserModel{}), errors.New("ошибка на стороне сервера")
	}
	return 
}

func UserUpdate(user models.UserModel) (err error) {
	if user.Password == ""{
		return errors.New("поля пароля не может быть пустым")
	}
	user.Password = utils.Heshing(user.Password)
	err = repository.UpdateUserToDB(user)
	if err != nil{
		return errors.New("ошибка на стороне сервера")
	}
	return
}


// Авторизация пользователя
func AuthorizationUserService(u models.UserModel) (token string,err error){
	
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
	token, err = middleware.GenerateToken(u.Login)
	if err != nil{
		return "", errors.New("ошибка при создании токена")
	}

	return token, nil
}


// Регистрация пользователя
func RegisterUserService(user models.UserModel) (token string,err error) {
	// Проверка наличия пользователья в Б.Д.
	inf, err := repository.GetUserFromDB(user.Login)
	if err != nil{
		return "", errors.New("ошибка")
	}
	if inf != (models.UserModel{}) {
		return "", errors.New("пользователь с таким логином уже зарегистрирован")
	}
	
	// Добавление нового пользоватля
	user.Password = utils.Heshing(user.Password)
	err = repository.CreateNewUserToDB(user)
	if err != nil{
		return "", errors.New("ошибка при создание нового пользователя")
	}
	
	// Добавление нового токена
	token, err = middleware.GenerateToken(user.Login)
	if err != nil{
		return "", errors.New("ошибка при создании токена")
	}

	return token, nil
}