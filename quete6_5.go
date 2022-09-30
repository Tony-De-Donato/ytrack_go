package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, lettre := range s {
		a := rune(lettre)
		z01.PrintRune(a)
	}
}

func main() {
	PrintStr("Hello World!")
}
