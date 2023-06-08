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
	r.Route("/departments", func(r chi.Router) {
		r.Get("/", handlers.GetAllDepartments)
		r.Post("/", handlers.CreateDepartment)
		r.Delete("/", handlers.DeleteDepartment)

		r.Get("/{id}", handlers.GetDepartmentByID)
		r.Put("/{id}", handlers.UpdateDepartment)
	})
	http.ListenAndServe(":8080", r)
}
