package main

import (
	"fmt"
)

func fibo(nb int) int {
	if nb < 0 {
		return -1
	}
	if nb <= 1 {
		return nb
	} else {
		return fibo(nb-1) + fibo(nb-2)
	}
}

func main() {
	arg := 8
	fmt.Println(fibo(arg))
}
