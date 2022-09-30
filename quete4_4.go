package main

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for _, element := range a {
		for i := 0; i < len(element); i++ {
			lettre := rune(element[i])
			z01.PrintRune(lettre)

		}
		z01.PrintRune('\n')
	}
}

func main() {
	a := []string{"Hello", "how", "are", "you?"}
	PrintWordsTables(a)

}
