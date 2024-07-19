package database

import (
	"database/sql"
	"log"
	"os"
	"strings"

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

    migrationFile := "/Users/apple/Desktop/Go-lang/Rescounts_Task/migrations/init.sql"

    content, err := os.ReadFile(migrationFile)
    if err != nil {
        log.Fatalf("Error reading migration file %s: %v", migrationFile, err)
    }

    // Split the content by semicolon to separate SQL statements
    sqlStatements := strings.Split(string(content), ";")

    // Execute each SQL statement
    for _, stmt := range sqlStatements {
        stmt = strings.TrimSpace(stmt)
        if stmt == "" {
            continue
        }
        _, err = DB.Exec(stmt)
        if err != nil {
            log.Fatalf("Error executing migration statement: %v\nStatement: %s", err, stmt)
        }
    }

}
