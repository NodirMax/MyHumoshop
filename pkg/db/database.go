package db

import (
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
	d := DB_struct{DB_user: "postgres", DB_password: 1010, DB_name: "humoshop"}
	conStr := fmt.Sprintf("user=%s password=%d dbname=%s sslmode=disable", d.DB_user, d.DB_password, d.DB_name)
	DB, err = sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal("database connection ERROR")
	}
}

// Закрыть Базу Данных
func DatabaseClose() {
	DB.Close()
}
