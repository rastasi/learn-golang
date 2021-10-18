package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/domain"
	"github.com/rastasi/learn-golang/crud/http/controller"
	"github.com/rastasi/learn-golang/crud/http/router"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func StartHttpServer(domain domain.Domain) {
	r := mux.NewRouter()
	r.StrictSlash(false)
	utils.AddRouter(r, "/albums", *router.AlbumRouter{
		AlbumController: controller.AlbumController{
			AlbumService: domain.AlbumService,
		},
	}.Init())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
