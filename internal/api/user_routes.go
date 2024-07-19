package api

import (
    "github.com/gorilla/mux"
    "Rescounts_Task/internal/handlers/user"
)

func RegisterUserRoutes(router *mux.Router) {
    router.HandleFunc("/signup", user.SignUpUser).Methods("POST")
    router.HandleFunc("/login", user.LoginUser).Methods("POST")
    router.HandleFunc("/credit-cards", user.AddCreditCard).Methods("POST")
    router.HandleFunc("/credit-cards/{id}", user.DeleteCreditCard).Methods("DELETE")
    router.HandleFunc("/products", user.ListProducts).Methods("GET")
    router.HandleFunc("/buy", user.BuyProducts).Methods("POST")
    router.HandleFunc("/orders", user.GetUserOrders).Methods("GET")
}
