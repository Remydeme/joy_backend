package handler

import (
	"encoding/json"
	"net/http"

	"github.com/joy/server/model"
)

const (
	passwordHashCost = 8
)

func Login(w http.ResponseWriter, r *http.Request) {

	var credential models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// session open here

	// hashe the password
	/*
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credential.Password), passwordHashCost)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err := database.DB.AddUser(credential.Username, hashedPassword); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}*/
}
