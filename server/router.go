package server

import (
	"net/http"

	"errors"

	"github.com/gorilla/mux"
	"github.com/joy/handler"
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
		Routes["signUp"] = Route{Name: "SignUp", Path: "/sign-up", Method: "Post", Handler: handler.SignUp}
		Routes["create-user"] = Route{Name: "create-user", Path: "/create-user", Method: "Post", Handler: handler.CreateUser}
		Routes["update-user"] = Route{Name: "update-user", Path: "/update-user", Method: "Post", Handler: handler.UpdateUser}
		Routes["get-user"] = Route{Name: "get-user", Path: "/get-user", Method: "Post", Handler: handler.GetUser}
		Routes["create-profile"] = Route{Name: "create-profile", Path: "/create-profile", Method: "Post", Handler: handler.CreateProfile}
		Routes["update-profile"] = Route{Name: "update-profile", Path: "/update-profile", Method: "Post", Handler: handler.UpdateProfile}
		Routes["get-profile"] = Route{Name: "get-profile", Path: "/get-profile", Method: "Post", Handler: handler.GetProfile}
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
