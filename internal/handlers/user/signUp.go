package user

import (
	"Rescounts_Task/internal/database"
	"Rescounts_Task/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(w http.ResponseWriter, r *http.Request) {
    var req models.SignUpRequest

    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    // Log the request body
    fmt.Println("Request Body:", string(body))

    // Decode the request body into the req struct
    err = json.Unmarshal(body, &req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // Create a new user ID
    userID := uuid.New()

    // Insert the new user into the database
    _, err = database.DB.Exec(
        "INSERT INTO users (id, username, email, password, role,stripe_customer_id) VALUES ($1, $2, $3, $4, $5, $6)",
        userID, req.Username, req.Email, string(hashedPassword), "user","",
    )
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Username or email already exists", http.StatusConflict)
        } else {
            fmt.Println("ERROR:", err)
            http.Error(w, "Error creating user", http.StatusInternalServerError)
        }
        return
    }

    // Respond with status 201 Created
    w.WriteHeader(http.StatusCreated)
}
