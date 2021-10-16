package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/http/router"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func StartHttpServer() {
	r := mux.NewRouter()
	r.StrictSlash(false)
	utils.AddRouter(r, "/albums", *router.AlbumRouter())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
