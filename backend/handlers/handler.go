package handlers

import (
    "net/http"
)

// JSONMiddleware sets the Content-Type header to application/json
func JSONMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

// ErrorHandler is an example of a general error handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
    w.WriteHeader(status)
    w.Write([]byte(message))
}
