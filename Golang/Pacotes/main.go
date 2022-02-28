package main

import (
	"fmt"
	submodulo "modulo/SubModulo"
)

func main() {
	fmt.Println("Mensagem em arquivo principal")
	submodulo.Escrever()
}
