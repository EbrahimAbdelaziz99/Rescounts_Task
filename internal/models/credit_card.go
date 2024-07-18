package models

import (
	"time"

	"github.com/google/uuid"
)

type CreditCard struct {
    ID           uuid.UUID `json:"id"`
    UserID       uuid.UUID `json:"user_id"`
    StripeCardID string    `json:"stripe_card_id"`
    LastFour     string    `json:"last_four"`
    Brand        string    `json:"brand"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type AddCreditCardRequest struct {
    UserID    uuid.UUID `json:"user_id"`
    CardToken string    `json:"card_token"`
}
