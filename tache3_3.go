package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	renvoi := rune(0)
	if n > 0 {
		if n < len(s)-1 {
			renvoi = rune(s[n-1])
		}
	} else if n < 0 {
		if -n < len(s)-1 {
			renvoi = rune(s[len(s)+(n)])
		}
	} else {
		renvoi = 0
	}
	return renvoi
}

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
