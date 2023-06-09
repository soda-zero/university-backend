package handlers

import (
	"encoding/json"
	"net/http"
	"zeroCalSoda/university-backend/private/db"
	"zeroCalSoda/university-backend/private/db/models"

	"github.com/go-chi/chi/v5"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	repo, err := db.NewCourseRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	courses, err := repo.GetCourses()
	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(courses)
	if err != nil {
		http.Error(w, "Failed to encode courses", http.StatusInternalServerError)
		return
	}
}
func GetCourseByID(w http.ResponseWriter, r *http.Request) {
	courseID := chi.URLParam(r, "id")

	repo, err := db.NewCourseRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	course, err := repo.GetCourseByID(courseID)
	if err != nil {
		http.Error(w, "Failed to fetch course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		http.Error(w, "Failed to encode course", http.StatusInternalServerError)
		return
	}
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course

	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Failed to decode course data", http.StatusBadRequest)
		return
	}

	repo, err := db.NewCourseRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.CreateCourse(course)
	if err != nil {
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	courseID := chi.URLParam(r, "id")

	repo, err := db.NewCourseRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.DeleteCourse(courseID)
	if err != nil {
		http.Error(w, "Failed to delete course", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	courseID := chi.URLParam(r, "id")
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Failed to decode course data", http.StatusBadRequest)
		return
	}

	repo, err := db.NewCourseRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.UpdateCourse(courseID, course)
	if err != nil {
		http.Error(w, "Failed to update course", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
