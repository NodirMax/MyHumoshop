package middleware

import (
	"HumoSHOP/internal/repository"
	"errors"
)

func AdminCheck(login string) (err error) {
	user, err := repository.GetUserFromDB(login)
	if err != nil{
		return errors.New("ошибка на стороне")
	}
	if user.Role != "admin"{
		return errors.New("не админ")
	}
	return nil
}