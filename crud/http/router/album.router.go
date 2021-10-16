package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/http/controller"
)

func AlbumRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetAlbums).Methods("GET")
	r.HandleFunc("/", controller.PostAlbum).Methods("POST")
	return r
}
