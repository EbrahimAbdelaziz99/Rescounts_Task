package admin

import (
	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"
	"encoding/json"
	// "log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var req models.CreateProductRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    // log.Fatalln(err)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    productID := uuid.New()
    _, err = database.DB.Exec(
        "INSERT INTO products (id, name, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
        productID, req.Name, req.Description, req.Price, time.Now(), time.Now(),
    )
    if err != nil {
        http.Error(w, "Error creating product", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
