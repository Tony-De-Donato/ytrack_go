package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	array := []string{}
	mot := ""
	for i := 0; i < len(s); i++ {
		mot += string(s[i])
		if s[i] == sep[0] {
			test := true
			for j := 0; j < len(sep); j++ {
				if s[i+j] != sep[j] {
					test = false
					mot += string(s[i])
				}
				if test {
					array = append(array, mot)
					mot = ""

				}
			}
		} else if i == len(s)-1 {
			array = append(array, mot)
			mot = ""
		}
	}
	return array
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
