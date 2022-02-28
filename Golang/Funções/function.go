package main

import "fmt"

func somar(x, y int8) int8 {
	return x + y
}

func calculos(x, y uint8) (uint8, uint8) {
	soma := x + y
	multi := x * y

	return soma, multi
}

func main() {

	fmt.Println(somar(10, 90))

	f := func(texto string) string {
		return texto
	}

	fmt.Println(f("Banana"))

	fmt.Println(calculos(80, 70))

	Resultadosoma, Resultadomulti := calculos(90, 90)
	fmt.Println(Resultadosoma, Resultadomulti)

	Resultadosoma1, _ := calculos(90, 90)
	fmt.Println(Resultadosoma1)
}
