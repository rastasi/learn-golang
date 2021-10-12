package main

import "fmt"

func pointerTest(szoveg *string) {
	*szoveg = "okok"
}

func nonPointerTest(szoveg string) string {
	szoveg = "okok2"
	return szoveg
}

func main() {
	var eredeti_szoveg = "OK"
	pointerTest(&eredeti_szoveg)
	fmt.Println(eredeti_szoveg)
	eredeti_szoveg = nonPointerTest(eredeti_szoveg)
	fmt.Println(eredeti_szoveg)
}
