package main

import "fmt"

func createMatrix() {
	a := [][]string{
		{"line", "tiger"},
		{"cat", "dog"},
		{"pigeon", "hamster"},
	}
	printSlices(a)
}

func printSlices(a [][]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	createMatrix()
}
