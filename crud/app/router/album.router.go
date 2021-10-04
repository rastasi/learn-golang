package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/app/controller"
)

func AlbumRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetAlbums).Methods("GET")
	r.HandleFunc("/", controller.PostAlbums).Methods("POST")
	return r
}
