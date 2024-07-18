package api

import (
	"Rescounts_Task/internal/handlers/admin"

	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(router *mux.Router) {
    s := router.PathPrefix("/admin").Subrouter();
    s.HandleFunc("/products", admin.CreateProduct).Methods("POST")
    s.HandleFunc("/products/{id}", admin.UpdateProduct).Methods("PUT")
    s.HandleFunc("/products/{id}", admin.DeleteProduct).Methods("DELETE")
    s.HandleFunc("/sales", admin.GetProductSales).Methods("POST")
}
