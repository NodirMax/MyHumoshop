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

// роут относящийся к категориям
	router3 := router.PathPrefix("/category").Subrouter()
	// public роуты
	router3.HandleFunc("", handlers.CategoryGET).Methods("GET")
	router3.HandleFunc("/{category_name}", handlers.CategoryGET_id).Methods("GET")
	// private роуты
	router3.Use(middleware.ProtectedEndpoint)
	router3.HandleFunc("", handlers.CategoryPOST_id).Methods("POST")
	router3.HandleFunc("", handlers.CategoryPUT_id).Methods("PUT")
	router3.HandleFunc("", handlers.CategoryDELETE_id).Methods("DELETE")

// роут относящийся к продукту
	router4 := router.PathPrefix("/product").Subrouter()
	// public роуты
	router4.HandleFunc("", handlers.ProductGet).Methods("GET")
	// private роуты
	router4.Use(middleware.ProtectedEndpoint)
	router4.HandleFunc("", handlers.ProductPOST).Methods("POST")
	router4.HandleFunc("", handlers.ProductPUT).Methods("PUT")
	router4.HandleFunc("", handlers.ProductDELETE).Methods("Delete")


	//Запуск сервера на порту 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER listing ERROR")
	}
}