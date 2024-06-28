package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/harshit-1245/auth/backend/models"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error loading .env file")
        os.Exit(1)
    }

    // Set CGO_ENABLED environment variable
    os.Setenv("CGO_ENABLED", os.Getenv("CGO_ENABLED"))

    uri := os.Getenv("TURSO_DATABASE_URL")
    if uri == "" {
        fmt.Fprintf(os.Stderr, "DB_URI environment variable not set")
        os.Exit(1)
    }

    db, err := sql.Open("sqlite3", uri)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to open db: %v", err)
        os.Exit(1)
    }

    DB = db
    fmt.Println("Database connection established")

    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    `
    _, err = DB.Exec(createTableSQL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create table: %v", err)
        os.Exit(1)
    }
    fmt.Println("Users table created or already exists")
}

func CloseDB() {
    if DB != nil {
        DB.Close()
        fmt.Println("Database connection closed")
    }
}

func CreateUser(db *sql.DB, user *models.User) error {
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
    return err
}

func FindUserByUsername(db *sql.DB, username string) (*models.User, error) {
    var user models.User
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
