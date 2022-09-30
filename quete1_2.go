package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 'z'; compteur >= 'a'; compteur-- {
		z01.PrintRune(compteur)
	}
}
