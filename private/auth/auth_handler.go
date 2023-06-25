package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var userData Users
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Failed to decode user", http.StatusInternalServerError)
		return
	}

	repo, err := NewAuthRepository()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	err = repo.Signup(userData)
	if err != nil {
		http.Error(w, "Failed to signup", http.StatusInternalServerError)
		return
	}
	fmt.Println(userData)
	w.WriteHeader(http.StatusCreated)
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var userData Users
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo, err := NewAuthRepository()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer repo.Close()

	sessionData, err := repo.Signin(userData.Username)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionData.Token,
		Expires: sessionData.Expiration,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Welcome back!",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
