package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/http/controller"
)

type AlbumRouter struct {
	AlbumController controller.AlbumControllerInterface
}

func (r AlbumRouter) Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", r.AlbumController.Index).Methods("GET")
	router.HandleFunc("/", r.AlbumController.Create).Methods("POST")
	return router
}
