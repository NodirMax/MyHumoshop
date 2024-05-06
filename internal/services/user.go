package services

import (
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