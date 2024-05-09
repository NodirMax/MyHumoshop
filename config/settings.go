package config

import (
	"encoding/json"
	"log"
	"os"
)

type configModel struct {
	ServerPort string           `json:"server_port"`
	PgSettings postgresSettings `json:"pg_settings"`
	JWTSecret  string           `json:"jwt_secret"`
}

type postgresSettings struct {
	User     string `json:"user"`
	DBName   string `json:"dbname"`
	Password int64 `json:"password"`
}

var Settings configModel

func ReadConfig(filePath string) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading config file!", err)
		return
	}

	err = json.Unmarshal(f, &Settings)
	if err != nil {
		log.Println("Error marshalling config file!", err)
		return
	}

}
