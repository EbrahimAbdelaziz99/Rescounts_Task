package user

import (
    "encoding/json"
    "log"
    "Rescounts_Task/internal/database"
    "Rescounts_Task/internal/models/order"
    "net/http"
    "time"

    "github.com/google/uuid"
)




func GetUserOrders(res http.ResponseWriter, req *http.Request) {
    userID := req.URL.Query().Get("user_id")
    if userID == "" {
        http.Error(res, "Missing user_id parameter", http.StatusBadRequest)
        return
    }

    rows, err := database.DB.Query("SELECT id, total_amount, created_at FROM orders WHERE user_id = $1", userID)
    if err != nil {
        http.Error(res, "Error retrieving orders", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        err := rows.Scan(&order.ID, &order.TotalAmount, &order.CreatedAt)
        if err != nil {
            http.Error(res, "Error scanning order", http.StatusInternalServerError)
            return
        }
        orders = append(orders, order)
    }

    json.NewEncoder(res).Encode(orders)
}
