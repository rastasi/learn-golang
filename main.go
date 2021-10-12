package main

import "fmt"

func pointertest(szoveg *string) {
	*szoveg = "okok"
}

func nempointertest(szoveg string) string {
	szoveg = "okok2"
	return szoveg
}

func mateMatrix() {
	a := [][]string{
		{"line", "tiger"},
		{"cat", "dog"},
		{"pigeon", "hamster"},
	}
	PrintSlices(a)
}

func PrintSlices(a [][]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

type Cat struct {
	name string
	age  int8
}

func tamasMap() {
	m := make(map[string]Cat)

	m["k1"] = Cat{
		name: "Cirmos",
		age:  10,
	}

	var masodik_cica Cat = Cat{
		name: "Khatgo",
		age:  2,
	}

	macska_varo(&masodik_cica)

	m["k2"] = masodik_cica

	for i, elem := range m {
		fmt.Println(i)
		fmt.Println(elem.name)
	}
}

func macska_varo(egyed *Cat) {
	egyed.name = "Atneveztunk"
	*egyed = Cat{}
}

func main() {
	var eredeti_szoveg = "OK"
	pointertest(&eredeti_szoveg)
	// fmt.Println(eredeti_szoveg)
	// eredeti_szoveg = nempointertest(eredeti_szoveg)
	// fmt.Println(eredeti_szoveg)
	// mateMatrix()
	tamasMap()
}
