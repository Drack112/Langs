package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!")
	escrever("Olá Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}
}
