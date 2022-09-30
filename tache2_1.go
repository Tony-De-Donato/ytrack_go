package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	somme := 1
	if nb >= 1 {
		for i := 1; i <= nb; i++ {
			somme = somme * i
		}
	} else {
		somme = 0
	}
	return somme

}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
