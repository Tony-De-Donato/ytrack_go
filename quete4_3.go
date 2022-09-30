package main

import (
	"fmt"
)

func ConcatParams(args []string) string {
	var reponse string
	for _, element := range args {
		reponse += element + "\n"
	}
	return reponse
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
