package user

import (
    // "log"
    "Rescounts_Task/internal/database"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    // "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/client"
)

func DeleteCreditCard(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cardID := vars["id"]

    sc := &client.API{}
    sc.Init(os.Getenv("STRIPE_SECRET_KEY"), nil)

    var stripeCardID string
    err := database.DB.QueryRow("SELECT stripe_card_id FROM credit_cards WHERE id = $1", cardID).Scan(&stripeCardID)
    if err != nil {
        http.Error(w, "Error finding credit card", http.StatusInternalServerError)
        return
    }

    _, err = sc.Cards.Del(stripeCardID, nil)
    if err != nil {
        http.Error(w, "Error deleting credit card", http.StatusInternalServerError)
        return
    }

    _, err = database.DB.Exec("DELETE FROM credit_cards WHERE id = $1", cardID)
    if err != nil {
        http.Error(w, "Error deleting credit card from database", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
