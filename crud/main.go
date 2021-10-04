package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud/app/router"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/albums", router.AlbumRouter())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
