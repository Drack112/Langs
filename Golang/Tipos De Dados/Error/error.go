package main

import (
	"errors"
	"fmt"
)

func main() {
	var erro error
	fmt.Println(erro)

	erro2 := errors.New("Erro Encontrado")
	fmt.Println(erro2)
}
