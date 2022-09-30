package main

import (
	"fmt"
)

func PrintNbrInOrder(nbr int) {
	longueur := string(nbr)
	if string(longueur[0]) == "-" {
		fmt.Printf("-")
	}
	minimum := int(longueur[0])
	for i := 0; i < len(longueur); i++ {
		if int(longueur[i]) <= minimum {
			minimum = int(longueur[i])
		}
	}

}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(231)
}
