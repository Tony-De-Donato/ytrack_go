package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	compteur := 0
	for i := 0; i < len(s)-1; i++ {
		if (s[i] < 91 && s[i] > 64) || (s[i] < 123 && s[i] > 96) {
			compteur += 1
		}
	}
	return compteur

}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
