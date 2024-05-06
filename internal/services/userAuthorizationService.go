package services

import (
	"HumoSHOP/internal/middleware"
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

func AuthorizationUserService(u models.UserModel) (token string,err error){
	
	inf, err := repository.CheckingUser(u.Login)
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