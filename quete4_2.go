package main

import (
	"fmt"
)

func MakeRange(min, max int) []int {
	if min >= max {
		var array []int
		return array
	} else {
		array2 := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			array2[i] = min + i
		}
		return array2
	}
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
