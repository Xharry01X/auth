package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() error {
    url := os.Getenv("TURSO_DATABASE_URL")
    authToken := os.Getenv("TURSO_AUTH_TOKEN")
    dsn := fmt.Sprintf("%s?authToken=%s", url, authToken)

    var err error
    DB, err = sql.Open("libsql", dsn)
    if err != nil {
        return fmt.Errorf("failed to open db: %w", err)
    }

    // Create schema
    err = createSchema(DB)
    if err != nil {
        return fmt.Errorf("failed to create schema: %w", err)
    }

    return nil
}

func createSchema(db *sql.DB) error {
    schema := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    `
    _, err := db.Exec(schema)
    if err != nil {
        log.Fatalf("Error creating schema: %v", err)
    }
    return err
}
