package main

import (
	"fmt"
)

func UltimateDivMod(a *int, b *int) {
	valeur_b := *b
	*b = *a % *b
	*a = (*a - *b) / valeur_b

}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
