package main

import (
	"github.com/rastasi/learn-golang/crud/http"
	"github.com/rastasi/learn-golang/crud/lib/database"
)

func main() {
	database.Setup()
	http.StartHttpServer()
}
