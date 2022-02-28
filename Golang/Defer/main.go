package main

import "fmt"

func funcao() {
    fmt.Println("Executando a função 1")
}

func funcao2() {
    fmt.Println("Executando a função 2")
}

func main() {
    // adiar a execução, deixar pra ultimo
    defer funcao()
    funcao2()
}
