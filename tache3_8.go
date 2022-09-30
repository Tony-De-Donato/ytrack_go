package main

import (
	"fmt"
)

func IsUppercase(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] < 64 || s[i] > 90 {
			return false

		}

	}
	return true
}
func main() {
	fmt.Println(IsUppercase("HELLO"))
	fmt.Println(IsUppercase("HELLO!"))

}
