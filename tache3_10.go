package main

import "fmt"

func IsAlpha(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] < 65 || s[i] > 91 {
			if s[i] < 96 || s[i] > 123 {
				return false
			}
		}

	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
