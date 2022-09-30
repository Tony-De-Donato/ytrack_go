package main

import (
	"fmt"
)

func IsSorted(f func(a, b int) int, a []int) bool {
	reponse := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) != -1 {
			reponse = false
		}
	}
	return reponse
}

func f(a, b int) int {
	switch {
	case a > b:
		return 1
	case a == b:
		return 0
	default:
		return -1
	}
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
