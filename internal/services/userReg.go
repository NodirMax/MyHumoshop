package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"errors"
)

func RegisterUserService(user models.UserModule) (err error) {
	// Проверка наличия пользователья в Б.Д.
	inf, err := repository.GetUserFromDB(user.Login)
	if err != nil{
		return errors.New("ошибка")
	}
	if inf != (models.UserModule{}) {
		return errors.New("пользователь с таким логином уже зарегистрирован")
	}
	
	// Добавление нового пользоватля
	err = repository.CreateNewUserToDB(user)
	if err != nil{
		return errors.New("ошибка при создание нового пользователя")
	}
	return nil
}