package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return nb
	} else {
		return nb * RecursiveFactorial(nb-1)
	}

}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
