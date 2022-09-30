package main

import (
	"fmt"
)

func sqrt(nb int) int {
	resultat := 0
	for i := 0; i < nb; i++ {
		if i*i == nb {
			resultat = i
		}
	}
	return resultat
}

func sqrt2(nb float64) float64 {
	// oui je sais que je suis trés fort en math ☺
	vmax := nb
	vmin := float64(0)
	variable := float64(0)
	for {
		variable = (vmax + vmin) / 2
		if variable*variable > nb {
			vmax = variable

		} else {
			vmin = variable

		}
		stock := variable*variable - nb
		if float64(float64(stock)*float64(stock)) <= float64(0.0000001) {
			break
		}

	}
	return variable

}

func main() {
	arg := 4
	fmt.Println(sqrt(arg))
}
