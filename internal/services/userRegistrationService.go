package services

import (
	"HumoSHOP/internal/middleware"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

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