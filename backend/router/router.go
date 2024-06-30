package router

import (
    "github.com/gorilla/mux"
    "github.com/harshit-1245/auth/backend/handlers"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
    router.Use(handlers.JSONMiddleware) // Apply the JSON middleware to all routes
}
