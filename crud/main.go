package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/app/router"
	"github.com/rastasi/learn-golang/crud/lib/database"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func startHttpServer() {
	r := mux.NewRouter()
	r.StrictSlash(false)
	utils.AddRouter(r, "/albums", *router.AlbumRouter())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	database.Setup()
	startHttpServer()
}
