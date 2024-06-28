package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harshit-1245/auth/backend/database"
	"github.com/harshit-1245/auth/backend/router"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	r := mux.NewRouter()
	router.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
