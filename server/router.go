package server

import (
	"net/http"

	"errors"

	"github.com/gorilla/mux"
	"github.com/joy/server/handler"
)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var Routes = make(map[string]Route)

var (
	ErrRoutesAllocFailed = errors.New("Allocation of map [string] Route failed")
)

func init() {
	if Routes != nil {
		Routes["login"] = Route{Name: "Login", Path: "/login", Method: "Post", Handler: handler.Login}
		Routes["register-user"] = Route{Name: "RegisterUser", Path: "/register-user/{id:[A-Z0-9]+-*}", Method: "Post", Handler: handler.RegisterUser}
		Routes["register-office"] = Route{Name: "RegisterOffice", Path: "/register-office", Method: "Post", Handler: handler.RegisterOffice}
		Routes["create-user"] = Route{Name: "CreateUser", Path: "/create-user", Method: "Post", Handler: handler.CreateUser}
		Routes["update-user"] = Route{Name: "UpdateUser", Path: "/update-user", Method: "Post", Handler: handler.UpdateUser}
		Routes["get-user"] = Route{Name: "GetUser", Path: "/get-user", Method: "Post", Handler: handler.GetUser}
		Routes["create-profile"] = Route{Name: "CreateProfile", Path: "/create-profile", Method: "Post", Handler: handler.CreateProfile}
		Routes["update-profile"] = Route{Name: "UpdateProfile", Path: "/update-profile", Method: "Post", Handler: handler.UpdateProfile}
		Routes["get-profile"] = Route{Name: "GetProfile", Path: "/get-profile", Method: "Post", Handler: handler.GetProfile}
	}
}

func NewRouter() (*mux.Router, error) {
	if Routes == nil {
		return nil, ErrRoutesAllocFailed
	}
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		handler := http.Handler(route.Handler)
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router, nil
}
