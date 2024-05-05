package myproject

import (
	"HumoSHOP/internal/handlers"
	"HumoSHOP/internal/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.AuthorizationUzer).Methods("POST")

	// роут относящийся к пользователю
	router2 := router.PathPrefix("/profile").Subrouter()
	// Используем middleware 
	router2.Use(middleware.ProtectedEndpoint)
	router2.HandleFunc("/", handlers.UserGET).Methods("GET")
	router2.HandleFunc("/", handlers.UserPUT).Methods("PUT")

	router3 := router.PathPrefix("/category").Subrouter()
	router3.HandleFunc("", handlers.CategoryGET).Methods("GET")
	router3.HandleFunc("/{category_name}", handlers.CategoryGET_id).Methods("GET")


	//Запуск сервера на порту 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER listing ERROR")
	}
}