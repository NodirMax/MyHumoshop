package main

import "HumoSHOP/pkg/db"

func main() {
	db.DatabaseConnect() // Соединение с базой данных
	defer db.DatabaseClose() // Закрытие соединение с базой данных
}