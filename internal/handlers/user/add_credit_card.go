package user

import (
	"encoding/json"
	// "log"
	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

func AddCreditCard(w http.ResponseWriter, r *http.Request) {
    var req models.AddCreditCardRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    sc := &client.API{}
    sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)

    card, err := sc.Cards.New(&stripe.CardParams{
        Customer: stripe.String(req.UserID.String()),
        Token:    stripe.String(req.CardToken),
    })
    if err != nil {
        http.Error(w, "Error adding credit card", http.StatusInternalServerError)
        return
    }

    _, err = database.DB.Exec(
        "INSERT INTO credit_cards (id, user_id, stripe_card_id, last_four, brand, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
        uuid.New(), req.UserID, card.ID, card.Last4, card.Brand, time.Now(), time.Now(),
    )
    if err != nil {
        http.Error(w, "Error saving credit card", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
