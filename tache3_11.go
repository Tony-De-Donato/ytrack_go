package main

import (
	"fmt"
)

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] < 48 || s[i] > 57 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
