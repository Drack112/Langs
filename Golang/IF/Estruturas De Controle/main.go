package main

import "fmt"

func main() {
	n1 := 10

	if n1 > 15 {
		fmt.Println("Maior do que 15")
	} else {
		fmt.Println("Menor do que 15")
	}

	if outronumero := n1; outronumero > 0 {
		fmt.Println("Número é maior")
	} else if n1 < -10 {
		fmt.Println("Entre 0 e -10")
	} else {
		fmt.Println("oof")
	}
}
