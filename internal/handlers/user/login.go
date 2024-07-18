package user

import (
    "database/sql"
    "encoding/json"
    // "log"
    "Rescounts_Task/internal/database"
    "Rescounts_Task/internal/models"
    "net/http"

    "golang.org/x/crypto/bcrypt"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var req models.LoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    var storedPassword string
    err = database.DB.QueryRow("SELECT password FROM users WHERE email = $1", req.Email).Scan(&storedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        } else {
            http.Error(w, "Error logging in user", http.StatusInternalServerError)
        }
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
}
