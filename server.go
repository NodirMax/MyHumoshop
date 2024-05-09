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
	routerProfile := router.PathPrefix("/profile").Subrouter()
	//private rout, Используем middleware 
	routerProfile.Use(middleware.ProtectedEndpoint)
	routerProfile.HandleFunc("/", handlers.UserGET).Methods("GET")
	routerProfile.HandleFunc("/", handlers.UserUpdate).Methods("PUT")

// роут относящийся к категориям
	routerCategory := router.PathPrefix("/category").Subrouter()
	// public роуты
	router.HandleFunc("/category", handlers.CategoryGET).Methods("GET")
	router.HandleFunc("/category/{category_name}", handlers.CategoryGETbyid).Methods("GET")
	// private роуты, Используем middleware 
	routerCategory.Use(middleware.ProtectedEndpoint)
	routerCategory.HandleFunc("/", handlers.CategoryCreate).Methods("POST")
	routerCategory.HandleFunc("/", handlers.CategoryUpdate).Methods("PUT")
	routerCategory.HandleFunc("/", handlers.CategoryDELETE).Methods("DELETE")

// роут относящийся к продукту
	routerProduct := router.PathPrefix("/product").Subrouter()
	// public роуты
	router.HandleFunc("/product", handlers.ProductGet).Methods("GET")
	// private роуты, Используем middleware 
	routerProduct.Use(middleware.ProtectedEndpoint)
	routerProduct.HandleFunc("/", handlers.ProductCreate).Methods("POST")
	routerProduct.HandleFunc("/", handlers.ProductUpdate).Methods("PUT")
	routerProduct.HandleFunc("/", handlers.ProductDELETE).Methods("Delete")

// роут относящийся к заказам
	routerOrder := router.PathPrefix("/order").Subrouter()
	// private роуты, Используем middleware
	routerOrder.Use(middleware.ProtectedEndpoint)
	routerOrder.HandleFunc("/", handlers.OrderCreate).Methods("POST")
	routerOrder.HandleFunc("/", handlers.OrderGet).Methods("GET")


	//Запуск сервера на порту 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER listing ERROR")
	}
}	