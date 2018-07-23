package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joy/server/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credential models.User
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// session open here
	log.Println(credential)
}
