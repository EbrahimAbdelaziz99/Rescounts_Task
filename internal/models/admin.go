package models

import (
    "github.com/google/uuid"
)

type ProductSalesRequest struct {
    FromDate string `json:"from_date"`
    ToDate   string `json:"to_date"`
    UserName string `json:"user_name"`
}

type ProductSalesResponse struct {
    ProductID uuid.UUID `json:"product_id"`
    Quantity  int       `json:"quantity"`
    Total     float64   `json:"total"`
}
