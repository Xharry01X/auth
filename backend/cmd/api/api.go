package api

import (
	"database/sql"
	"fmt"
	"github.com/go-fuego/fuego "
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// APIServer struct holds the address of the server and the database connection
type APIServer struct {
	addr string
	db   *sql.DB
}

// NewAPIServer is a constructor function that initializes and returns an instance of APIServer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the HTTP server and listens for incoming requests
func (s *APIServer) Run() error {
	// Create a new Fuego server
	server := fuego.NewServer(nil)

	// HandleFunc sets up a handler for the /ping endpoint
	server.HandleFunc("/ping", func(w fuego.ResponseWriter, r *fuego.Request) {
		// Respond with "pong" when /ping is accessed
		w.Write([]byte("pong"))
	})

	// Log the address where the server is starting
	fmt.Printf("Starting server at %s\n", s.addr)
	// Start the HTTP server and listen on the specified address
	return server.ListenAndServe(s.addr)
}
