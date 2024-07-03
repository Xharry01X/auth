package main

import (
	"log"
	"net/http"

	
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/harshit-1245/auth/backend/database"
	"github.com/harshit-1245/auth/backend/router"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize the database
    err = database.InitDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    defer database.DB.Close()

    log.Println("Database connected successfully")

    // Setup router and routes
    r := mux.NewRouter()
    router.RegisterRoutes(r)



    // Configure CORS
    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Start the server
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}
