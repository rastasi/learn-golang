package main_test

import (
	"fmt"
	"testing"

	"github.com/rastasi/learn-golang/crud-learn-2/controller"
	"github.com/rastasi/learn-golang/crud-learn-2/repository"
)

func ControllerTest(t *testing.T) {
	personController := controller.PersonController{
		PersonRepository: repository.PersonRepositoryMock{},
	}
	fmt.Println(personController)
}
