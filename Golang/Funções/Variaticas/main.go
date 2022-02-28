package main

import "fmt"

func somar(n ...int) int {
    total := 0
    for _, numeros := range n {
        total += numeros
    }
    return total
}

func escrever(texto string, numeros ...int) {
    for _, numeros := range numeros {
        fmt.Println(texto, numeros)
    }
}

func main() {
    fmt.Println(somar(80, 90, 471))
    escrever("arroz", 9)
}
