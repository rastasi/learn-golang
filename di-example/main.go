package main

import (
	"fmt"

	"github.com/rastasi/learn-golang/di-example/controller"
	"github.com/rastasi/learn-golang/di-example/database"
)

func main() {
	c := controller.Controller{
		Database: database.Database{},
	}
	response := c.Index()
	fmt.Println(response)
}
