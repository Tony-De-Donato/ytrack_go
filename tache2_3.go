package main

import "fmt"

func IterativePower(nb int, power int) int {
	resultat := 1
	for i := 0; i < power; i++ {
		resultat = resultat * nb
	}
	return resultat
}

func main() {
	fmt.Println(IterativePower(5, 2))
}
