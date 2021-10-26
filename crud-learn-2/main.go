package main

import (
	"net/http"

	"github.com/rastasi/learn-golang/crud-learn-2/controller"
	"github.com/rastasi/learn-golang/crud-learn-2/repository"
	"github.com/rastasi/learn-golang/crud-learn-2/router"
)

func muxServer() {
	personRepository := repository.PersonRepository{}
	personController := controller.PersonController{
		PersonRepository: personRepository,
	}
	personRouter := router.PersonRouter(personController)

	http.ListenAndServe(":8090", personRouter)
}

// CREATE - CREATE
// READ   - INDEX
// READ   - SHOW
// UPDATE - UPDATE
// DELETE - DESTROY

func main() {
	repository.ConnectDb()
	muxServer()
}
