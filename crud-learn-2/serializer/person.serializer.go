package serializer

import "github.com/rastasi/learn-golang/crud-learn-2/model"

type PersonSerializer struct{}

type PersonResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Gender string `json:"gender"`
}

func (s PersonSerializer) Serialize(person model.Person) PersonResponse {
	var gender string

	if person.Gender == 1 {
		gender = "male"
	} else {
		gender = "female"
	}

	return PersonResponse{
		ID:     person.ID,
		Name:   person.Name,
		Age:    person.Age,
		Gender: gender,
	}
}
