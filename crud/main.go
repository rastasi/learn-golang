package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/app/router"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func main() {
	r := mux.NewRouter()
	utils.AddRouter(r, "/albums", *router.AlbumRouter())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
