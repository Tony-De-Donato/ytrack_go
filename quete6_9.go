package main

import (
	"fmt"
	"strconv"
)

func BasicAtoi(s string) int {
	var output int

	stringLen := len(s)

	for i := 0; i < stringLen; i++ {
		tempNum, _ := strconv.Atoi(string(s[i]))
		output = output*10 + tempNum
	}
	return output
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
