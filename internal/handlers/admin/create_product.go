package admin

import (
    "encoding/json"
    "log"
    "Rescounts_Task/internal/database"
    "net/http"
    "time"

    "github.com/google/uuid"
)

type CreateProductRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var req CreateProductRequest
    err := json.NewDecoder(r.Body).Decode(&req)
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
