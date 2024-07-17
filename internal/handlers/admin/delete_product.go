package admin

import (
    "log"
    "Rescounts_Task/internal/database"
    "net/http"

    "github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    productID := vars["id"]

    _, err := database.DB.Exec("DELETE FROM products WHERE id = $1", productID)
    if err != nil {
        http.Error(w, "Error deleting product", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
