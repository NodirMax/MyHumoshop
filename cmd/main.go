package main

import (
	"HumoSHOP/api"
	"HumoSHOP/config"
	"HumoSHOP/pkg/db"
)

func main() {
	config.ReadConfig("config/config.json")

	
	db.DatabaseConnect()// Соединение с базой данных
	defer db.DatabaseClose()// Закрытие соединение с базой данных
	api.StartRouter()
}