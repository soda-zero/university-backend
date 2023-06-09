package handlers

import (
	"encoding/json"
	"net/http"
	"zeroCalSoda/university-backend/private/db/models"
	"zeroCalSoda/university-backend/private/db/repositories"

	"github.com/go-chi/chi/v5"
)

func GetCareerLevels(w http.ResponseWriter, r *http.Request) {
	repo, err := repositories.NewCareerLevelRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	careerLevels, err := repo.GetCareerLevels()
	if err != nil {
		http.Error(w, "Failed to fetch career career levels", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(careerLevels)
	if err != nil {
		http.Error(w, "Failed to encode career level", http.StatusInternalServerError)
		return
	}
}
func GetCareerLevelByID(w http.ResponseWriter, r *http.Request) {
	careerLevelID := chi.URLParam(r, "id")

	repo, err := repositories.NewCareerLevelRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	careerLevel, err := repo.GetCareerLevelByID(careerLevelID)
	if err != nil {
		http.Error(w, "Failed to fetch career level", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(careerLevel)
	if err != nil {
		http.Error(w, "Failed to encode career level", http.StatusInternalServerError)
		return
	}
}

func CreateCareerLevel(w http.ResponseWriter, r *http.Request) {
	var careerLevel models.CareerLevel

	err := json.NewDecoder(r.Body).Decode(&careerLevel)
	if err != nil {
		http.Error(w, "Failed to decode career level data", http.StatusBadRequest)
		return
	}

	repo, err := repositories.NewCareerLevelRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.CreateCareerLevel(careerLevel.Name)
	if err != nil {
		http.Error(w, "Failed to create career level", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteCareerLevel(w http.ResponseWriter, r *http.Request) {
	careerLevelID := chi.URLParam(r, "id")

	repo, err := repositories.NewCareerLevelRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.DeleteCareerLevel(careerLevelID)
	if err != nil {
		http.Error(w, "Failed to delete career level", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateCareerLevel(w http.ResponseWriter, r *http.Request) {
	careerLevelID := chi.URLParam(r, "id")
	var careerLevel models.CareerLevel
	err := json.NewDecoder(r.Body).Decode(&careerLevel)
	if err != nil {
		http.Error(w, "Failed to decode career level data", http.StatusBadRequest)
		return
	}

	repo, err := repositories.NewCareerLevelRepository()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.UpdateCareerLevel(careerLevelID, careerLevel.Name)
	if err != nil {
		http.Error(w, "Failed to update career level", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
