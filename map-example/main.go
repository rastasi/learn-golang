package main

import "fmt"

type Cat struct {
	name string
	age  int8
}

func createMap() {
	m := make(map[string]Cat)

	m["k1"] = Cat{
		name: "Cirmos",
		age:  10,
	}

	m["k2"] = Cat{
		name: "Karthago",
		age:  2,
	}

	for i, elem := range m {
		fmt.Println(i)
		fmt.Println(elem.name)
	}
}

func main() {
	createMap()
}
