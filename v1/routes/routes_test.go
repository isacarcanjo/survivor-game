package routes

import (
	config2 "serve/v1/config"
	"testing"

	"github.com/gorilla/mux"
)

func makeSut() *mux.Router {
	r := mux.NewRouter()
	return r
}

func TestRouters(t *testing.T) {
	router := makeSut()
	t.Parallel()
	config := config2.GetDotEnv()

	_, err := CreateRoutes(router, GetSurvivorRouters(), config)

	if err != nil {
		t.Error(err)
	}
}
