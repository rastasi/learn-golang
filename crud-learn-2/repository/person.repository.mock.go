package repository

import "github.com/rastasi/learn-golang/crud-learn-2/model"

type PersonRepositoryMock struct {
}

func (r PersonRepositoryMock) All() []model.Person {
	return []model.Person{}
}

func (r PersonRepositoryMock) FindById(personId int) model.Person {
	return model.Person{}
}

func (r PersonRepositoryMock) Create(person *model.Person) {
}

func (r PersonRepositoryMock) Update(person *model.Person) {
}

func (r PersonRepositoryMock) Destroy(personId int) {
}
