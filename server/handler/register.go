package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/joy/server/database"

	"github.com/gorilla/mux"
	"github.com/joy/server/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	log.Println(errors.New("Remy is here"))

	params := mux.Vars(r)

	// parse json
	var office models.Office

	err := json.NewDecoder(r.Body).Decode(&office)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte("Error while trying to read "))

		return
	}

	log.Println("id", params["id"], "   body:", office)

	tx := database.Database.Begin()

	tx.Insert(office)

	err = tx.Commit()

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte("Error while trying to read "))

		return

	}
}

func RegisterOffice(w http.ResponseWriter, r *http.Request) {
}
