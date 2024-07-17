package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error
    connStr := os.Getenv("DATABASE_URL")
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
}
