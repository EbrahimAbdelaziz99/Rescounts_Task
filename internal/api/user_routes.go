package api

import (
	"Rescounts_Task/internal/handlers/user"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
    s := router.PathPrefix("/users").Subrouter();
    s.HandleFunc("/signup", user.SignUpUser).Methods("POST")
    s.HandleFunc("/login", user.LoginUser).Methods("POST")
    s.HandleFunc("/credit-cards", user.AddCreditCard).Methods("POST")
    s.HandleFunc("/credit-cards/{id}", user.DeleteCreditCard).Methods("DELETE")
    s.HandleFunc("/products", user.ListProducts).Methods("GET")
    s.HandleFunc("/buy", user.BuyProducts).Methods("POST")
    s.HandleFunc("/orders", user.GetUserOrders).Methods("GET")
}
