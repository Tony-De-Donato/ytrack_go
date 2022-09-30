package main

import (
	"fmt"
)

func IsPrime(nb int) bool {
	resultat := true
	for i := 2; i < nb; i++ {
		for j := 2; j < nb; j++ {
			if i*j == nb {
				resultat = false
			}
		}

	}
	return resultat

}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
