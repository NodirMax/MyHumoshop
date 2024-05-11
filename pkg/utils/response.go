package utils

import (
	"encoding/json"
	"log"
)

func ResponseJson(text string) (resp []byte) {
	resp, err := json.Marshal(text)
	if err != nil{
		log.Println("[INFO]: json marshal ERROR")
	}
	return
}