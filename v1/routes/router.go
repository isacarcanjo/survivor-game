package routes

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"serve/v1/config"
	. "serve/v1/helpers"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

var methods = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE"}

type Router struct {
	URI                   string
	Method                string
	DomainFunc            func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func CreateRoutes(r *mux.Router, routes []Router, config config.Config) (*mux.Router, error) {
	api := r.PathPrefix(fmt.Sprintf("/%s", config.Version)).Subrouter()
	api.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	//api.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("swagger"))))
	for _, route := range routes {
		if len(route.URI) == 0 {
			return nil, errors.New("URI is empty in method " + route.Method)
		}
		if reflect.TypeOf(route.DomainFunc).Kind() != reflect.Func {
			return nil, errors.New("DomainFunc to be nil in method " + route.Method)
		}
		if !Contains(methods, route.Method) {
			return nil, errors.New("Methods is wrong in method " + route.Method)
		}
		api.HandleFunc(route.URI, route.DomainFunc).Methods(route.Method)
	}
	return r, nil
}
