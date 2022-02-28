package main

import "fmt"

func main() {
	var variavel1 string = "Hello World"
	// ou
	variavel2 := "Hello World" // var variavel 2 string = "Hello World"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "OOf"
		variavel4 string = "Oof"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Uma var", "Outra Var"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Uma constante" // Não pode ser alterada

	fmt.Println(constante1)
}
