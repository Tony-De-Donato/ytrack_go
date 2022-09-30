package main

import "fmt"

func Index(s string, toFind string) int {
	reponse := -1
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			for j := 0; j < len(toFind); j++ {
				test := true
				if s[i+j] != toFind[j] {
					test = false
				}
				if test {
					reponse = i
				}
			}
		}
	}
	return reponse
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
