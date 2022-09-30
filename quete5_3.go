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

func Any(f func(string) bool, a []string) bool {
	rep := false
	for i := 0; i < len(a); i++ {
		if f(string(a[i])) {
			rep = true
		}
	}
	return rep
}

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
