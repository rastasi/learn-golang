package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type KiscicaResponse struct {
	Name string
	Age  int8
}

func kiscicaHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World"))
	response, _ := json.Marshal(KiscicaResponse{
		Name: "Cirmi",
		Age:  2,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/kiscica/", kiscicaHandler).Methods("GET")
	http.ListenAndServe(":8090", router)
}
