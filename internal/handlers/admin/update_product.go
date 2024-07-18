package admin

import (
	"encoding/json"
	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    productID := vars["id"]

    var req models.UpdateProductRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec(
        "UPDATE products SET name = $1, description = $2, price = $3, updated_at = $4 WHERE id = $5",
        req.Name, req.Description, req.Price, time.Now(), productID,
    )
    if err != nil {
        http.Error(w, "Error updating product", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
