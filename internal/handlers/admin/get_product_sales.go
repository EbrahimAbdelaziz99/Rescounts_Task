package admin

import (
	"database/sql"
	"encoding/json"

	// "log"
	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"
	"net/http"
	"time"
)

func GetProductSales(w http.ResponseWriter, r *http.Request) {
    var req models.ProductSalesRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    var rows *sql.Rows
    var query string
    var args []interface{}

    query = `SELECT p.id, sum(oi.quantity) as quantity, sum(oi.quantity * oi.price) as total 
             FROM products p 
             JOIN order_items oi ON p.id = oi.product_id 
             JOIN orders o ON oi.order_id = o.id 
             JOIN users u ON o.user_id = u.id`

    if req.FromDate != "" && req.ToDate != "" {
        from, fromDateError := time.Parse(time.RFC3339, req.FromDate)
        if fromDateError != nil {
            http.Error(w, "Invalid FromDate format. Expected format: RFC3339", http.StatusBadRequest)
            return
        }

        to, toDateError := time.Parse(time.RFC3339, req.ToDate)
        if toDateError != nil {
            http.Error(w, "Invalid ToDate format. Expected format: RFC3339", http.StatusBadRequest)
            return
        }

        query += " WHERE o.created_at BETWEEN $1 AND $2"
        args = append(args, from, to)
    }

    if req.UserName != "" {
        if len(args) > 0 {
            query += " AND u.username = $3"
            args = append(args, req.UserName)
        } else {
            query += " WHERE u.username = $1"
            args = append(args, req.UserName)
        }
    }

    query += " GROUP BY p.id"

    rows, err = database.DB.Query(query, args...)
    if err != nil {
        http.Error(w, "Error retrieving product sales", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var sales []models.ProductSalesResponse
    for rows.Next() {
        var s models.ProductSalesResponse
        err := rows.Scan(&s.ProductID, &s.Quantity, &s.Total)
        if err != nil {
            http.Error(w, "Error scanning sales data", http.StatusInternalServerError)
            return
        }
        sales = append(sales, s)
    }

    json.NewEncoder(w).Encode(sales)
}
