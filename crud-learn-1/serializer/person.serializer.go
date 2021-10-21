package serializer

import "github.com/rastasi/learn-golang/crud-learn-1/model"

type PersonSerializer struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Gender string `json:"gender"`
}

func Serialize(person model.Person) PersonSerializer {
	var gender string

	if person.Gender == 1 {
		gender = "male"
	} else {
		gender = "female"
	}

	return PersonSerializer{
		ID:     person.ID,
		Name:   person.Name,
		Age:    person.Age,
		Gender: gender,
	}
}
