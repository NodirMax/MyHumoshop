package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Пароль, который нужно захешировать
func Heshing(password string) (p string){

	// Генерируем хеш пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка при хешировании пароля:", err)
		return
	}
	return string(hashedPassword)
}

func HeshChecking(hashedPassword string ,password string)(err error){
    // Проверяем пароль
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return
    }

    return nil
}