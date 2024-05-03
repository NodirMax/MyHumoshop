package services

import (
	"HumoSHOP/internal/middleware"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

func RegisterUserService(user models.UserModels) (token string,err error) {
	// Проверка наличия пользователья в Б.Д.
	inf, err := repository.GetUserFromDB(user.Login)
	if err != nil{
		return "", errors.New("ошибка")
	}
	if inf != (models.UserModels{}) {
		return "", errors.New("пользователь с таким логином уже зарегистрирован")
	}
	
	// Добавление нового пользоватля
	user.Password = utils.Heshing(user.Password)
	err = repository.CreateNewUserToDB(user)
	if err != nil{
		return "", errors.New("ошибка при создание нового пользователя")
	}
	
	// Добавление нового токена
	token = middleware.CreateToken(user.Login)
	err = repository.CreateTokenDB(user.Id, token)
	if err != nil{
		return "",errors.New("ошибка")
	}

	return token, nil
}