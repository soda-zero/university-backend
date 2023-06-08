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

		r.Get("/{id}", handlers.GetDepartmentByID)
		r.Put("/{id}", handlers.UpdateDepartment)
		r.Delete("/{id}", handlers.DeleteDepartment)
	})

	r.Route("/career-levels", func(r chi.Router) {
		r.Get("/", handlers.GetCareerLevels)
		r.Post("/", handlers.CreateCareerLevel)

		r.Get("/{id}", handlers.GetCareerLevelByID)
		r.Put("/{id}", handlers.UpdateCareerLevel)
		r.Delete("/{id}", handlers.DeleteCareerLevel)
	})
	http.ListenAndServe(":8080", r)
}
