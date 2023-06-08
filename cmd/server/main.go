package main

import (
	"net/http"
	"zeroCalSoda/university-backend/private/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.HelloHandler)
	r.Get("/departments", handlers.GetAllDepartments)
	r.Get("/departments/{id}", handlers.GetAllDepartmentByID)
	http.ListenAndServe(":8080", r)
}
