package handlers

import (
	"net/http"
)

func UserGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SECRET PAGE"))
}