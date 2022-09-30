package main

import (
	"fmt"
)

func ToLower(s string) string {
	resultat := ""

	for i := 0; i < len(s); i++ {
		if s[i] < 91 && s[i] > 64 {
			resultat += string(rune(int(s[i]) + 32))
		} else {
			resultat += string(s[i])
		}
	}
	return resultat
}

func ToUpper(s string) string {
	resultat := ""

	for i := 0; i < len(s); i++ {
		if s[i] < 123 && s[i] > 96 {
			resultat += string(rune(int(s[i]) - 32))
		} else {
			resultat += string(s[i])
		}
	}
	return resultat
}

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

func Capitalize(s string) string {
	resultat := ToUpper(string(s[0]))
	for i := 1; i < len(s); i++ {
		if string(s[i]) == " " || string(s[i]) == "+" {
			resultat += string(s[i]) + ToUpper(string(s[i+1]))
			i += 1
		} else {
			resultat += ToLower(string(s[i]))
		}

	}
	return resultat
}

func main() {
	fmt.Println(Capitalize("hello! how are you? How+Are+Things+4you?"))
}
