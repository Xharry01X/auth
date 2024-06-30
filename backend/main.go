package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/harshit-1245/auth/backend/database"
    "github.com/harshit-1245/auth/backend/router"
)

func main() {
    err := database.InitDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    defer database.DB.Close()

    r := mux.NewRouter()
    router.RegisterRoutes(r)

    log.Fatal(http.ListenAndServe(":8080", r))
}
