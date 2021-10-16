package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/http/controller"
	"github.com/rastasi/learn-golang/crud/http/router"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func StartHttpServer() {
	r := mux.NewRouter()
	r.StrictSlash(false)
	utils.AddRouter(r, "/albums", *router.AlbumRouter{
		AlbumController: controller.AlbumController{
			AlbumService: service.AlbumService{
				AlbumRepository: repository.AlbumRepository{},
			},
		},
	}.Init())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
