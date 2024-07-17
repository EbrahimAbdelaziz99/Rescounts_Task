package user

import (
    "encoding/json"
    "log"
    "Rescounts_Task/internal/database"
    "net/http"
    "time"

    "github.com/google/uuid"
)

type BuyProductsRequest struct {
    UserID      uuid.UUID `json:"user_id"`
    Products    []Product `json:"products"`
}

type Product struct {
    ProductID uuid.UUID `json:"product_id"`
    Quantity  int       `json:"quantity"`
}

func BuyProducts(w http.ResponseWriter, r *http.Request) {
    var req BuyProductsRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    var totalAmount float64
    for _, p := range req.Products {
        var price float64
        err := database.DB.QueryRow("SELECT price FROM products WHERE id = $1", p.ProductID).Scan(&price)
        if err != nil {
            http.Error(w, "Error finding product price", http.StatusInternalServerError)
            return
        }
        totalAmount += price * float64(p.Quantity)
    }

    orderID := uuid.New()
    _, err = database.DB.Exec("INSERT INTO orders (id, user_id, total_amount, created_at) VALUES ($1, $2, $3, $4)", orderID, req.UserID, totalAmount, time.Now())
    if err != nil {
        http.Error(w, "Error creating order", http.StatusInternalServerError)
        return
    }

    for _, p := range req.Products {
        _, err := database.DB.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price, created_at) VALUES ($1, $2, $3, $4, $5, $6)", uuid.New(), orderID, p.ProductID, p.Quantity, price, time.Now())
        if err != nil {
            http.Error(w, "Error creating order items", http.StatusInternalServerError)
            return
        }
    }

    w.WriteHeader(http.StatusCreated)
}
