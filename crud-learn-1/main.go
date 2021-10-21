package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud-learn-1/controller"
	"github.com/rastasi/learn-golang/crud-learn-1/repository"
)

func muxServer() {
	router := mux.NewRouter()

	router.HandleFunc("/people/", controller.Index).Methods("GET")
	router.HandleFunc("/people/", controller.Create).Methods("POST")
	router.HandleFunc("/people/{personId}", controller.Show).Methods("GET")
	router.HandleFunc("/people/{personId}", controller.Update).Methods("PUT")
	router.HandleFunc("/people/{personId}", controller.Destroy).Methods("DELETE")

	http.ListenAndServe(":8090", router)

}

// CREATE - CREANTE
// READ   - INDEX
// READ   - SHOW
// UPDATE - UPDATE
// DELETE - DESTROY

func main() {
	repository.ConnectDb()
	muxServer()
}
