package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joy/server/database"

	"github.com/gorilla/mux"
	"github.com/joy/server/model"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

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

	database.DB.Add(office)

}

func RegisterOffice(w http.ResponseWriter, r *http.Request) {
}
