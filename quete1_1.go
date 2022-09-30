package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 'a'; compteur <= 'z'; compteur++ {
		z01.PrintRune(compteur)
	}
}
