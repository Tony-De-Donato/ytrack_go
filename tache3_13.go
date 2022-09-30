package main

import (
	"fmt"
)

func ToUpper(s string) string {
	resultat := ""

	for i := 0; i < len(s); i++ {
		if s[i] < 123 && s[i] > 96 {
			resultat += string(rune(int(s[i]) - 32))
		} else {
			resultat += string(s[i])
		}
	}
	return resultat
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
