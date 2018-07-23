package server

import (
	"net/http"

	_ "github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

const (
	port = ":8080"
)

func init() {
}

func initXORM() {

}

func (a *App) initApp() error {
	var err error
	a.Router, err = NewRouter()
	return err
}

func (a *App) Run() (bool, error) {
	a.initCache()
	if err := a.initApp(); err != nil {
		return false, err
	}

	if err := http.ListenAndServe(port, a.Router); err != nil {
		return false, err
	}

	return true, nil
}
