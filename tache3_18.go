package main

import "fmt"

func Join(elems []string, spr string) string {
	resultat := elems[0]
	for i := 1; i < len(elems); i++ {
		resultat += spr + elems[i]
	}
	return resultat
}

func main() {
	txt := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(txt, ":"))
}
