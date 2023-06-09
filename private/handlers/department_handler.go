package handlers

import (
	"encoding/json"
	"net/http"
	"zeroCalSoda/university-backend/private/db/models"
	"zeroCalSoda/university-backend/private/db/repositories"

	"github.com/go-chi/chi/v5"
)

func GetDepartments(w http.ResponseWriter, r *http.Request) {
	repo, err := repositories.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	departments, err := repo.GetDepartments()
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
func GetDepartmentByID(w http.ResponseWriter, r *http.Request) {
	departmentID := chi.URLParam(r, "id")

	repo, err := repositories.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	department, err := repo.GetDepartmentByID(departmentID)
	if err != nil {
		http.Error(w, "Failed to fetch department", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(department)
	if err != nil {
		http.Error(w, "Failed to encode department", http.StatusInternalServerError)
		return
	}
}

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var department models.Department

	err := json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		http.Error(w, "Failed to decode department data", http.StatusBadRequest)
		return
	}

	repo, err := repositories.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.CreateDepartment(department.Name)
	if err != nil {
		http.Error(w, "Failed to create department", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	departmentID := chi.URLParam(r, "id")
	repo, err := repositories.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.DeleteDepartment(departmentID)
	if err != nil {
		http.Error(w, "Failed to delete department", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	departmentID := chi.URLParam(r, "id")
	var department models.Department
	err := json.NewDecoder(r.Body).Decode(&department)
	if err != nil {
		http.Error(w, "Failed to decode department data", http.StatusBadRequest)
		return
	}

	repo, err := repositories.NewDepartmentRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.UpdateDepartment(departmentID, department.Name)
	if err != nil {
		http.Error(w, "Failed to update department", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
