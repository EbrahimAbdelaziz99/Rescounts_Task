package api

import (
	"Rescounts_Task/internal/handlers/admin"

	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(router *mux.Router) {
    router.HandleFunc("/products", admin.CreateProduct).Methods("POST")
    router.HandleFunc("/products/{id}", admin.UpdateProduct).Methods("PUT")
    router.HandleFunc("/products/{id}", admin.DeleteProduct).Methods("DELETE")
    router.HandleFunc("/sales", admin.GetProductSales).Methods("POST")
}
