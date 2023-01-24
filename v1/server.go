package server

import (
	"fmt"
	"log"
	"net/http"
	. "serve/v1/config"
	"serve/v1/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Server interface {
	Start()
}

type HTTP struct {
	Config
}

func (s HTTP) Start() {
	r := mux.NewRouter()

	routesList := routes.GetRoutes()
	res, err2 := routes.CreateRoutes(r, routesList, s.Config)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("Listerning microservice users with PORT:%d in %s \n", s.Port, s.Env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), res))
}
