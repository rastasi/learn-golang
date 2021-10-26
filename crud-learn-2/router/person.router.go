package router

import (
	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud-learn-2/controller"
)

func PersonRouter(personController controller.PersonController) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/people/", personController.Index).Methods("GET")
	router.HandleFunc("/people/", personController.Create).Methods("POST")
	router.HandleFunc("/people/{personId}", personController.Show).Methods("GET")
	router.HandleFunc("/people/{personId}", personController.Update).Methods("PUT")
	router.HandleFunc("/people/{personId}", personController.Destroy).Methods("DELETE")
	return router
}
