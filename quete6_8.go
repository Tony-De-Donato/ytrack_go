package main

import (
	"fmt"
)

func StrRev(s string) string {
	str := ""
	for _, element := range s {
		str = string(element) + str
	}
	return str
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
