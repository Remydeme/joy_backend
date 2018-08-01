package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joy/server/database"
)

type App struct {
	Router *mux.Router
}

var DB *database.DB_s

const (
	user     = "wisseadmin"
	password = "Kaisershtul1996"
	host     = "localhost"
	port     = 5432
	dbname   = "wisse"
	sslmode  = "disable"
)

const (
	server_port = ":8080"
)

func init() {
}

func (a *App) initDB() error {
	DB = new(database.DB_s)
	err := DB.Connect_to_DB()
	return err
}

func (a *App) initApp() error {
	var err error
	a.Router, err = NewRouter()
	return err
}

func (a *App) Run() (bool, error) {

	if err := a.initDB(); err != nil {
		return false, err
	}

	if err := a.initApp(); err != nil {
		return false, err
	}

	if err := http.ListenAndServe(server_port, a.Router); err != nil {
		return false, err
	}

	return true, nil
}
