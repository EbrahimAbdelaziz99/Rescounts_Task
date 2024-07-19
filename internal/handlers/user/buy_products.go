package user

import (
	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"github.com/stripe/stripe-go/paymentintent"
)

func BuyProducts(w http.ResponseWriter, r *http.Request) {
    var req models.BuyProductsRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if len(req.Products) == 0 {
        http.Error(w, "No products to buy", http.StatusBadRequest)
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

    sc := &client.API{}
    sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)

    paymentIntentParams := &stripe.PaymentIntentParams{
        Amount: stripe.Int64(1099),
        Currency: stripe.String(string(stripe.CurrencyUSD)),
    };
    paymentIntent, err := paymentintent.New(paymentIntentParams);
    if err != nil {
        log.Printf("Error creating payment intent: %v", err)
        http.Error(w, "Error creating payment intent", http.StatusInternalServerError)
        return
    }

    if paymentIntent.Status != stripe.PaymentIntentStatusSucceeded {
        http.Error(w, "Payment not successful", http.StatusPaymentRequired)
        return
    }

    // Create order in database
    orderID := uuid.New()
    _, err = database.DB.Exec("INSERT INTO orders (id, user_id, total_amount, created_at) VALUES ($1, $2, $3, $4)", orderID, req.UserID, totalAmount, time.Now())
    if err != nil {
        log.Printf("Error creating order: %v", err)
        http.Error(w, "Error creating order", http.StatusInternalServerError)
        return
    }

    for _, p := range req.Products {
        _, err := database.DB.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price, created_at) VALUES ($1, $2, $3, $4, $5, $6)", uuid.New(), orderID, p.ProductID, p.Quantity, p.Price, time.Now())
        if err != nil {
            log.Printf("Error creating order items: %v", err)
            http.Error(w, "Error creating order items", http.StatusInternalServerError)
            return
        }
    }

    // Respond to the client
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Order created successfully",
        "order_id": orderID.String(),
    })
}
