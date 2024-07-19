package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

func AddCreditCard(w http.ResponseWriter, r *http.Request) {
	var req models.AddCreditCardRequest

	// Read and decode the request body into the req struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
        fmt.Println("ERROR:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize Stripe client
	sc := &client.API{}
	sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)

	// Check if user has a Stripe customer ID saved in the database
	var stripeCustomerID string
	err = database.DB.QueryRow("SELECT stripe_customer_id FROM users WHERE id = $1", req.UserID).Scan(&stripeCustomerID)
	if err == sql.ErrNoRows {
		// Create a new Stripe customer
		customerParams := &stripe.CustomerParams{
			Description: stripe.String("Customer for " + req.UserID.String()),
		}
		cust, err := sc.Customers.New(customerParams)
		if err != nil {
            fmt.Println("ERROR:", err)
			http.Error(w, "Error creating Stripe customer", http.StatusInternalServerError)
			return
		}
		stripeCustomerID = cust.ID

		// Save the Stripe customer ID in the database
		_, err = database.DB.Exec("UPDATE users SET stripe_customer_id = $1 WHERE id = $2", stripeCustomerID, req.UserID)
		if err != nil {
            fmt.Println("ERROR:", err)
			http.Error(w, "Error saving Stripe customer ID", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
        fmt.Println("ERROR:", err)
		http.Error(w, "Error retrieving Stripe customer ID", http.StatusInternalServerError)
		return
	}

	// Create a new card for the customer
	cardParams := &stripe.CardParams{
		Customer: stripe.String(stripeCustomerID),
		Token:    stripe.String(req.CardToken),
	}
	card, err := sc.Cards.New(cardParams)
	if err != nil {
        fmt.Println("ERROR:", err)
		http.Error(w, "Error adding credit card to Stripe", http.StatusInternalServerError)
		return
	}

	// Insert the credit card details into the database
	_, err = database.DB.Exec(
		"INSERT INTO credit_cards (id, user_id, stripe_card_id, last_four, brand, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		uuid.New(), req.UserID, card.ID, card.Last4, card.Brand, time.Now(), time.Now(),
	)
	if err != nil {
        fmt.Println("ERROR:", err)
		http.Error(w, "Error saving credit card to database", http.StatusInternalServerError)
		return
	}

	// Respond with a 201 Created status
	w.WriteHeader(http.StatusCreated)
}