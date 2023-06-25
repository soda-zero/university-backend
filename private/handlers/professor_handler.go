package handlers

import (
	"encoding/json"
	"net/http"
	"zeroCalSoda/university-backend/private/db/models"
	"zeroCalSoda/university-backend/private/db/repositories"

	"github.com/go-chi/chi/v5"
)

func GetProfessors(w http.ResponseWriter, r *http.Request) {
	repo, err := repositories.NewProfessorRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	professors, err := repo.GetProfessors()
	if err != nil {
		http.Error(w, "Failed to fetch professors", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(professors)
	if err != nil {
		http.Error(w, "Failed to encode professors", http.StatusInternalServerError)
		return
	}
}
func GetProfessorByID(w http.ResponseWriter, r *http.Request) {
	professorID := chi.URLParam(r, "id")
	repo, err := repositories.NewProfessorRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	professor, err := repo.GetProfessorByID(professorID)
	if err != nil {
		http.Error(w, "Failed to fetch professor", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(professor)
	if err != nil {
		http.Error(w, "Failed to encode professor", http.StatusInternalServerError)
		return
	}
}

func CreateProfessor(w http.ResponseWriter, r *http.Request) {
	var professorData models.Professor
	err := json.NewDecoder(r.Body).Decode(&professorData)
	if err != nil {
		http.Error(w, "Failed to decode professor", http.StatusInternalServerError)
		return
	}
	repo, err := repositories.NewProfessorRepository()
	if err != nil {
		http.Error(w, "Failed to connec to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.CreateProfessor(professorData)
	if err != nil {
		http.Error(w, "Failed to create professor", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	professorID := chi.URLParam(r, "id")
	repo, err := repositories.NewProfessorRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.DeleteProfessor(professorID)
	if err != nil {
		http.Error(w, "Failed to delete professor", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	professorID := chi.URLParam(r, "id")
	var professorData models.Professor
	err := json.NewDecoder(r.Body).Decode(&professorData)
	if err != nil {
		http.Error(w, "Failed to decode professor", http.StatusInternalServerError)
		return
	}

	repo, err := repositories.NewProfessorRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	err = repo.UpdateProfessor(professorID, professorData)
	if err != nil {
		http.Error(w, "Failed to update professor", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
