package db

import (
	"HumoSHOP/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB_struct struct {
	DB_user     string
	DB_password int
	DB_name     string
}

var DB *sql.DB

// Открыть Базу Данных
func DatabaseConnect() {
	var err error
	pgConf := config.Settings.PgSettings
	dsn := fmt.Sprintf("user=%s dbname=%s password=%d sslmode=disable", pgConf.User, pgConf.DBName, pgConf.Password)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("[INFO] Can't connect to Database!", err)
	}
}

// Закрыть Базу Данных
func DatabaseClose() {
	err := DB.Close()
	if err != nil {
		log.Println("[ERROR] Can't close Database connection!", err)
	}
}
