package main_test

import (
	"fmt"
	"testing"

	"github.com/rastasi/learn-golang/di-example/controller"
	"github.com/rastasi/learn-golang/di-example/database"
)

func TestController(t *testing.T) {
	// t.Skip()
	c := controller.Controller{
		Database: database.DatabaseMock{},
	}
	response := c.Index()
	fmt.Println(response)
	if response[0] != "mocked" {
		t.Error()
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
