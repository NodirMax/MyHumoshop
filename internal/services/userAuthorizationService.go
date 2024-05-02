package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

func AuthorizationUserService(u models.UserModels) (err error){
	// передаем хеширований пароль
	inf, err := repository.CheckingUser(u.Login, utils.Heshing(u.Password))
	if err != nil{
		return errors.New("ошибка на стороне сервера")
	}
	if inf.Id == 0{
		return errors.New("такого пользователя нет")
	}

	return nil
}