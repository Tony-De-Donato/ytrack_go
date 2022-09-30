package main

import "fmt"

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

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for i := nb; ; i++ {
			if IsPrime(i) {
				return i
			}

		}
	}

}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(20))
}
