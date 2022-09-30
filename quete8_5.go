package main

import (
	"os"

	"github.com/01-edu/z01"
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
	if os.Args[0] == "--upper" {
		for _, char := range os.Args[1:] {
			if len(char) == 1 {
				val := rune(char[0] - 49)
				val += 65
				z01.PrintRune(val)
			} else if len(char) == 2 {
				val2 := rune(char[0]*10 - 49 + char[1] + 97)
				z01.PrintRune(val2)
			}
		}
	} else {
		for _, char := range os.Args {
			if len(char) == 1 {
				val := rune(char[0] - 49)
				val += 97
				z01.PrintRune(val)
			} else if len(char) == 2 {
				val2 := rune(char[0]*10 - 49 + char[1] + 97 + 32)
				z01.PrintRune(val2)
			}
		}
	}
}
