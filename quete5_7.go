package main

import (
	"fmt"
)

func SortWordArr(a []string) {
	resultat := []string{}
	base := a
	for i := 0; i < len(a); i++ {
		if len(base) == 0 {
			break
		}
	}
	a = resultat
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
