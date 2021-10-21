package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud-learn-1/model"
	"github.com/rastasi/learn-golang/crud-learn-1/repository"
	"github.com/rastasi/learn-golang/crud-learn-1/serializer"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Index(w http.ResponseWriter, r *http.Request) {
	people := repository.All()

	var serialized []serializer.PersonSerializer

	for _, person := range people {
		serialized = append(serialized, serializer.Serialize(person))
	}

	respondWithJSON(w, http.StatusOK, serialized)
}

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personId, _ := strconv.Atoi(params["personId"])
	person := repository.FindById(personId)
	respondWithJSON(w, http.StatusOK, serializer.Serialize(person))
}

func Create(w http.ResponseWriter, r *http.Request) {
	var person model.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository.Create(&person)
	respondWithJSON(w, http.StatusCreated, serializer.Serialize(person))
}

func Update(w http.ResponseWriter, r *http.Request) {
	var body model.Person
	params := mux.Vars(r)
	personId, _ := strconv.Atoi(params["personId"])
	person := repository.FindById(personId)
	modified := person

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	modified.Name = body.Name
	modified.Age = body.Age
	modified.Gender = body.Gender

	repository.Update(&person)

	respondWithJSON(w, http.StatusOK, serializer.Serialize(person))
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personId, _ := strconv.Atoi(params["personId"])
	repository.Destroy(personId)
	w.WriteHeader(204)
}
