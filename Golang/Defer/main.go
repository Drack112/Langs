package main

import "fmt"

func funcao1() {
    fmt.Println("Hello Função 1")
}

func funcao2() {
    fmt.Println("Hello Função 2")
}

func main() {
    // adiar
    defer funcao1()
    funcao2()
}
