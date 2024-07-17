package main

import (
    "log"
    "Rescounts_Task/internal/database"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "Rescounts_Task/internal/api"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    database.InitDB()

    router := mux.NewRouter()
    api.RegisterRoutes(router)

    http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
