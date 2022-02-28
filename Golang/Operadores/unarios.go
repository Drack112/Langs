package main

import "fmt"

func main() {
	numero := 10
	numero++
	fmt.Println(numero) //11
	numero += numero
	fmt.Println(numero) //22

	numero--
	fmt.Println(numero) //21
	numero -= numero
	fmt.Println(numero) //0
}
