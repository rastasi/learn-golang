package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name string `json:"nev"`
	Age  int8   `json:"eletkor"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	person1 := Person{
		Name: "Tasi",
		Age:  34,
	}
	person2 := Person{
		Name: "Norbi",
		Age:  30,
	}
	respondWithJSON(w, http.StatusOK, []Person{person1, person2})
}

func muxServer() {
	router := mux.NewRouter()
	router.HandleFunc("/index", indexHandler).Methods("GET")
	http.ListenAndServe(":8090", router)
}

func main() {
	muxServer()
}
