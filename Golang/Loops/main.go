package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
	}

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	nome := [3]string{"João", "Diego", "Fernando"}

	for i, nome := range nome {
		fmt.Println(i, nome)
	}

	for _, nome := range nome {
		fmt.Println(nome)
	}

	for i, palavra := range "PALAVRA" {
		fmt.Println(i, palavra)
	}

	for i, palavra := range "PALAVRA" {
		fmt.Println(i, string(palavra))
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Carvalho",
	}

	for i, usuario := range usuario {
		fmt.Println(i, usuario)
	}

	o := 0
	for {
		o++
		fmt.Println(o)
	}
}
