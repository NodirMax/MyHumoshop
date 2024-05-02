package main

import (
	myproject "HumoSHOP"
	"HumoSHOP/pkg/db"
)

func main() {
	db.DatabaseConnect() // Соединение с базой данных
	defer db.DatabaseClose() // Закрытие соединение с базой данных

	myproject.StartRouter()
}