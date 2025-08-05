package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThuraMinThein/social-golang/internal/store"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user store.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := app.store.Users.Create(context.Background(), &user); err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (app *application) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := app.store.Users.GetAll()
	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}