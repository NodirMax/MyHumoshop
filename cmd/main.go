package main

import (
	myproject "HumoSHOP"
	"HumoSHOP/config"
	"HumoSHOP/pkg/db"
	"log"
)

func main() {
	log.Println("Starting to read config file...")
	config.ReadConfig("config/config.json")
	log.Println("Success!")

	
	db.DatabaseConnect()// Соединение с базой данных
	defer db.DatabaseClose()// Закрытие соединение с базой данных
	myproject.StartRouter()
}