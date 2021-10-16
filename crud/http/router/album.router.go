package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/http/controller"
)

func AlbumRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.AlbumController{}.Index).Methods("GET")
	r.HandleFunc("/", controller.AlbumController{}.Create).Methods("POST")
	return r
}
