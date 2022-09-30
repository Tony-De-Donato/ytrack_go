package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	donnee := os.Args
	var list []string
	for _, element := range donnee {
		list = append(list, string(element))
	}

	var list2 []string
	for len(list) != 0 {
		min := list[0]
		index_min := 0
		for i := 0; i < len(list); i++ {
			if list[i] < min {
				min = list[i]
				index_min = i
			}

		}
		list2 = append(list2, min)
		list = append(list[:index_min], list[index_min+1:]...)
	}
	for _, element := range list2 {
		z01.PrintRune(rune(element[0]))
		z01.PrintRune('\n')
	}
}
