package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/http/controller"
)

func AlbumRouter() *mux.Router {
	r := mux.NewRouter()
	c := controller.AlbumController{
		AlbumService: service.AlbumService{
			AlbumRepository: repository.AlbumRepository{},
		},
	}
	r.HandleFunc("/", c.Index).Methods("GET")
	r.HandleFunc("/", c.Create).Methods("POST")
	return r
}
