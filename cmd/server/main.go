package main

import (
	"net/http"
	"zeroCalSoda/university-backend/private/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", handlers.HelloHandler)
	http.ListenAndServe(":8080", r)
}
