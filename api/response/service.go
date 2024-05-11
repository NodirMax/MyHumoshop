package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorJsonMessage(w http.ResponseWriter, resp Resp) {
	w.Header().Set("Content-Type", "application/json")

	if resp.StatusCode == 0 {
		resp.StatusCode = 500
	}

	if resp.Message == "" {
		resp.Message = Errors[resp.StatusCode]
	}

	w.WriteHeader(resp.StatusCode)

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[ERROR] Can't Marshal Error JSON! Info: %v", err)
	}

	w.Write(respBytes)
}

func SuccessJsonMessage(w http.ResponseWriter, resp Resp) {
	w.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[ERROR] Can't Marshal Success JSON! Info: %v", err)
	}

	w.Write(respBytes)
}
