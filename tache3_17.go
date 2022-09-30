package main

import "fmt"

func BasicJoin(elems []string) string {
	resultat := ""
	for i := 0; i < len(elems); i++ {
		resultat += elems[i]
	}
	return resultat
}

func main() {
	txt := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(txt))
}
