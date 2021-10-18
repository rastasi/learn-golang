package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/http/controller"
)

type ArtistRouter struct {
	ArtistController controller.ArtistControllerInterface
}

func (r ArtistRouter) Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", r.ArtistController.Index).Methods("GET")
	router.HandleFunc("/", r.ArtistController.Create).Methods("POST")
	return router
}
