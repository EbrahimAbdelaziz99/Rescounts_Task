package api

import (
    "github.com/gorilla/mux"
    "Rescounts_Task/internal/handlers/admin"
)

func RegisterAdminRoutes(router *mux.Router) {
    router.HandleFunc("/admin/products", admin.CreateProduct).Methods("POST")
    router.HandleFunc("/admin/products/{id}", admin.UpdateProduct).Methods("PUT")
    router.HandleFunc("/admin/products/{id}", admin.DeleteProduct).Methods("DELETE")
    router.HandleFunc("/admin/sales", admin.GetProductSales).Methods("POST")
}
