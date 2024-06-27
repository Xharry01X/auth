package main

import (
	"log"
	"net/http"

	"github.com/harshit-1245/auth/backend/database"
	"github.com/harshit-1245/auth/backend/router"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	database.InitDB("mongodb://localhost:27017", "testdb")

	// Create a new mux router
	r := mux.NewRouter()

	// Register routes
	router.RegisterRoutes(r)

	// Start the server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
