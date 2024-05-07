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
	//private rout, Используем middleware 
	router2.Use(middleware.ProtectedEndpoint)
	router2.HandleFunc("/", handlers.UserGET).Methods("GET")
	router2.HandleFunc("/", handlers.UserUpdate).Methods("PUT")

// роут относящийся к категориям
	router3 := router.PathPrefix("/category").Subrouter()
	// public роуты
	router.HandleFunc("/category", handlers.CategoryGET).Methods("GET")
	router.HandleFunc("/category/{category_name}", handlers.CategoryGETbyid).Methods("GET")
	// private роуты, Используем middleware 
	router3.Use(middleware.ProtectedEndpoint)
	router3.HandleFunc("", handlers.CategoryCreate).Methods("POST")
	router3.HandleFunc("", handlers.CategoryUpdate).Methods("PUT")
	router3.HandleFunc("", handlers.CategoryDELETE).Methods("DELETE")

// роут относящийся к продукту
	router4 := router.PathPrefix("/product").Subrouter()
	// public роуты
	router.HandleFunc("/product", handlers.ProductGet).Methods("GET")
	// private роуты, Используем middleware 
	router4.Use(middleware.ProtectedEndpoint)
	router4.HandleFunc("", handlers.ProductCreate).Methods("POST")
	router4.HandleFunc("", handlers.ProductUpdate).Methods("PUT")
	router4.HandleFunc("", handlers.ProductDELETE).Methods("Delete")

// роут относящийся к заказам
	router5 := router.PathPrefix("/order").Subrouter()
	// private роуты, Используем middleware
	router5.Use(middleware.ProtectedEndpoint)
	router5.HandleFunc("", handlers.OrderCreate).Methods("POST")


	//Запуск сервера на порту 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER listing ERROR")
	}
}	