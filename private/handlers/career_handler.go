package handlers

import (
	"encoding/json"
	"net/http"
	"zeroCalSoda/university-backend/private/db/models"
	"zeroCalSoda/university-backend/private/db/repositories"

	"github.com/go-chi/chi/v5"
)

func GetCareers(w http.ResponseWriter, r *http.Request) {
	repo, err := repositories.NewCareerRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()
	careers, err := repo.GetCareers()
	if err != nil {
		http.Error(w, "Failed to fetch careers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(careers)
	if err != nil {
		http.Error(w, "Failed to encode careers", http.StatusInternalServerError)
		return
	}
}
func GetCareerByID(w http.ResponseWriter, r *http.Request) {
	careerID := chi.URLParam(r, "id")
	repo, err := repositories.NewCareerRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	career, err := repo.GetCareerByID(careerID)
	if err != nil {
		http.Error(w, "Failed to fetch career", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(career)
	if err != nil {
		http.Error(w, "Failed to fetch encode career", http.StatusInternalServerError)
		return
	}

}
func CreateCareer(w http.ResponseWriter, r *http.Request) {
	var careerData models.Career
	err := json.NewDecoder(r.Body).Decode(&careerData)
	if err != nil {
		http.Error(w, "Failed to decode career data", http.StatusBadRequest)
		return
	}
	repo, err := repositories.NewCareerRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.CreateCareer(careerData)
	if err != nil {
		http.Error(w, "Failed to create career", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteCareer(w http.ResponseWriter, r *http.Request) {
	careerID := chi.URLParam(r, "id")

	repo, err := repositories.NewCareerRepository()
	if err != nil {
		http.Error(w, "Failed to conenct to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.DeleteCareer(careerID)
	if err != nil {
		http.Error(w, "Failed to delete career", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateCareer(w http.ResponseWriter, r *http.Request) {
	careerID := chi.URLParam(r, "id")
	var careerData models.Career
	err := json.NewDecoder(r.Body).Decode(&careerData)
	if err != nil {
		http.Error(w, "Failed to decode career data", http.StatusInternalServerError)
		return
	}
	repo, err := repositories.NewCareerRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.UpdateCareer(careerID, careerData)
	if err != nil {
		http.Error(w, "Error updating career", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
