package main

import (
	"fmt"
)

func ToLower(s string) string {
	resultat := ""

	for i := 0; i < len(s); i++ {
		if s[i] < 91 && s[i] > 64 {
			resultat += string(rune(int(s[i]) + 32))
		} else {
			resultat += string(s[i])
		}
	}
	return resultat
}

func main() {
	fmt.Println(ToLower("HELLO! HOW ARE YOU?"))
}
