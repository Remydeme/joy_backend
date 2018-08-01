package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joy/server"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credential models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// session open here
	log.Println(credential)
	server.DB.AddUser(credential.Username, credential.Password)
}
