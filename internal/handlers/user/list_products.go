package user

import (
    "encoding/json"
    // "log"
    "Rescounts_Task/internal/database"
    "Rescounts_Task/internal/models"
    "net/http"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT id, name, description, price FROM products")
    if err != nil {
        http.Error(w, "Error retrieving products", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var p models.Product
        err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
        if err != nil {
            http.Error(w, "Error scanning product", http.StatusInternalServerError)
            return
        }
        products = append(products, p)
    }

    json.NewEncoder(w).Encode(products)
}
