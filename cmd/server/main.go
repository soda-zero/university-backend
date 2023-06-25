package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zeroCalSoda/university-backend/private/auth"
	"zeroCalSoda/university-backend/private/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	const PORT string = "8080"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", helloHandler)

	r.Route("/departments", func(r chi.Router) {
		r.Get("/", handlers.GetDepartments)
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
	r.Route("/courses", func(r chi.Router) {
		r.Get("/", handlers.GetCourses)
		r.Post("/", handlers.CreateCourse)

		r.Get("/{id}", handlers.GetCourseByID)
		r.Put("/{id}", handlers.UpdateCourse)
		r.Delete("/{id}", handlers.DeleteCourse)
	})

	r.Route("/careers", func(r chi.Router) {
		r.Get("/", handlers.GetCareers)
		r.Post("/", handlers.CreateCareer)

		r.Get("/{id}", handlers.GetCareerByID)
		r.Put("/{id}", handlers.UpdateCareer)
		r.Delete("/{id}", handlers.DeleteCareer)
	})
	r.Route("/professors", func(r chi.Router) {
		r.Get("/", handlers.GetProfessors)
		r.Post("/", handlers.CreateProfessor)

		r.Get("/{id}", handlers.GetProfessorByID)
		r.Put("/{id}", handlers.UpdateProfessor)
		r.Delete("/{id}", handlers.DeleteProfessor)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", auth.SignupHandler)
		r.Post("/signin", auth.SigninHandler)
	})
	fmt.Println("Listening on port: " + PORT)
	http.ListenAndServe(":"+PORT, r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello fella 🍌"
	response := map[string]interface{}{
		"message": message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
