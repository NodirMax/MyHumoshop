package services

import (
	"HumoSHOP/internal/models"
	"HumoSHOP/internal/repository"
	"HumoSHOP/pkg/utils"
	"errors"
)

func GetUserFromService(login string) ( user models.UserModels,err error) {
	user, err = repository.GetUserFromDB(login)
	if err != nil{
		return (models.UserModels{}), errors.New("ошибка на стороне сервера")
	}
	return 
}

func PutUserFromService(user models.UserModels) (err error) {
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