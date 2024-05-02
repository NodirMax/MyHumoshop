package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

func AuthorizationUserService(u models.UserModels) (err error){
	// передаем хеширований пароль
	inf, err := repository.CheckingUser(u.Login)
	if err != nil{
		return errors.New("ошибка на стороне сервера")
	}
	if inf == (models.UserModels{}){
		return errors.New("такого пользователя нет")
	}
	z :=  utils.HeshChecking(inf.Password, u.Password)
	if z != nil{
		return errors.New("такого пользователя нет")
	}
	
	return nil
}