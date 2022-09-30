package main

import (
	"fmt"
)

func Islowercase(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			return false
		}

	}
	return true
}
func main() {
	fmt.Println(Islowercase("hello"))
	fmt.Println(Islowercase("hello!"))

}
