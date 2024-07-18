package models

import (
    "github.com/google/uuid"
    "time"
)

type Order struct {
    ID          uuid.UUID `json:"id"`
    UserID      uuid.UUID `json:"user_id"`
    TotalAmount float64   `json:"total_amount"`
    CreatedAt   time.Time `json:"created_at"`
}

type OrderItem struct {
    ID        uuid.UUID `json:"id"`
    OrderID   uuid.UUID `json:"order_id"`
    ProductID uuid.UUID `json:"product_id"`
    Quantity  int       `json:"quantity"`
    Price     float64   `json:"price"`
    CreatedAt time.Time `json:"created_at"`
}

type BuyProductsRequest struct {
    UserID   uuid.UUID       `json:"user_id"`
    Products []OrderItem `json:"products"`
}
