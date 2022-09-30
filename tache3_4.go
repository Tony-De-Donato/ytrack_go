package main

import "fmt"

func Compare(a, b string) int {
	resultat := 0
	if len(a) == len(b) {
		resultat = 0
	} else if len(a) < len(b) {
		resultat = -1
	} else {
		resultat = 1
	}
	return resultat
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
