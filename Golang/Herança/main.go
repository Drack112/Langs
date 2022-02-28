package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Vitor", 19, 1.81}
	fmt.Println(p1)

	e1 := estudante{p1, "T.I", "USP"}
	fmt.Println(e1)

	fmt.Println(e1.pessoa.nome)
	fmt.Println(e1.nome)
	fmt.Println(e1.pessoa.sobrenome)
	fmt.Println(e1.pessoa)
	fmt.Println(e1.faculdade)
}
