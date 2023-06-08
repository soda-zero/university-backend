package handlers

import (
	"encoding/json"
	"net/http"
	"zeroCalSoda/university-backend/private/db"

	"github.com/go-chi/chi/v5"
)

func GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	repo, err := db.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	departments, err := repo.GetAllDepartments()
	if err != nil {
		http.Error(w, "Failed to fetch departments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(departments)
	if err != nil {
		http.Error(w, "Failed to encode departments", http.StatusInternalServerError)
		return
	}
}
func GetAllDepartmentByID(w http.ResponseWriter, r *http.Request) {
	departmentID := chi.URLParam(r, "id")

	repo, err := db.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	department, err := repo.GetDepartmentByID(departmentID)
	if err != nil {
		http.Error(w, "Failed to fetch departments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(department)
	if err != nil {
		http.Error(w, "Failed to encode departments", http.StatusInternalServerError)
		return
	}
}
