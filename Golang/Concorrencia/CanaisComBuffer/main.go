package main

import "fmt"

func main() {
	canal := make(chan string, 20)

	canal <- "Olá Mundo"
	canal <- "Olá Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
