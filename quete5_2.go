package main

import (
	"fmt"
)

func IsPrime(nb int) bool {
	resultat := true
	if nb < 2 {
		resultat = false
	}
	for i := 2; i < nb; i++ {
		for j := 2; j < nb; j++ {
			if i*j == nb {
				resultat = false
			}
		}

	}
	return resultat

}

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for _, c := range a {
		result = append(result, f(c))
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
