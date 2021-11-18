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

	switch person.Gender {
	case 1:
		gender = "male"
	case 2:
		gender = "female"
	case 3:
		gender = "dolphin"
	}

	return PersonResponse{
		ID:     person.ID,
		Name:   person.Name,
		Age:    person.Age,
		Gender: gender,
	}
}
