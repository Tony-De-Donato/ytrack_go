package main

import (
	"github.com/01-edu/z01"
)

func FirstRune(s string) rune {
	lettre := s[len(s)-1]
	val := rune(lettre)

	return val
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
