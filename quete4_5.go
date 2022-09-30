package main

import (
	"fmt"
)

func SplitWhiteSpaces(s string) []string {
	array := []string{}
	var mot string
	for i := 0; i < len(s); i++ {
		mot += string(s[i])
		if s[i] == 32 || i == len(s)-1 {
			array = append(array, mot)
			mot = ""
		}
	}
	return array
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
