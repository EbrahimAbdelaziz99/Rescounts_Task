package user

import (
    "database/sql"
    "encoding/json"
    "log"
    "Rescounts_Task/internal/database"
    "Rescounts_Task/internal/models"
    "net/http"

    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

func SignUpUser(w http.ResponseWriter, r *http.Request) {
    var req models.SignUpRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    userID := uuid.New()
    _, err = database.DB.Exec(
        "INSERT INTO users (id, username, email, password, role) VALUES ($1, $2, $3, $4, $5)",
        userID, req.Username, req.Email, string(hashedPassword), "user",
    )
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Username or email already exists", http.StatusConflict)
        } else {
            http.Error(w, "Error creating user", http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
}
