package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rastasi/learn-golang/crud-learn-2/model"
	"github.com/rastasi/learn-golang/crud-learn-2/repository"
	"github.com/rastasi/learn-golang/crud-learn-2/serializer"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type PersonController struct {
	PersonRepository repository.PersonRepositoryInterface
}

func (c PersonController) Index(w http.ResponseWriter, r *http.Request) {
	people := c.PersonRepository.All()

	var serialized []serializer.PersonResponse

	for _, person := range people {
		serialized = append(serialized, serializer.PersonSerializer{}.Serialize(person))
	}

	respondWithJSON(w, http.StatusOK, serialized)
}

func (c PersonController) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personId, _ := strconv.Atoi(params["personId"])
	person := c.PersonRepository.FindById(personId)
	respondWithJSON(w, http.StatusOK, serializer.PersonSerializer{}.Serialize(person))
}

func (c PersonController) Create(w http.ResponseWriter, r *http.Request) {
	var person model.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.PersonRepository.Create(&person)
	respondWithJSON(w, http.StatusCreated, serializer.PersonSerializer{}.Serialize(person))
}

func (c PersonController) Update(w http.ResponseWriter, r *http.Request) {
	var body model.Person
	params := mux.Vars(r)
	personId, _ := strconv.Atoi(params["personId"])
	person := c.PersonRepository.FindById(personId)
	modified := person

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	modified.Name = body.Name
	modified.Age = body.Age
	modified.Gender = body.Gender

	c.PersonRepository.Update(&person)

	respondWithJSON(w, http.StatusOK, serializer.PersonSerializer{}.Serialize(person))
}

func (c PersonController) Destroy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personId, _ := strconv.Atoi(params["personId"])
	c.PersonRepository.Destroy(personId)
	w.WriteHeader(204)
}
