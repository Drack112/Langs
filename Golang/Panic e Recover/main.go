package main

import "fmt"

func recuperar() {
    if r := recover(); r != nil {
        fmt.Println("Execução Recuperada")
    }
}

func alunoEstaAprovado(n1, n2 float32) bool {
    defer recuperar()
    media := (n1 + n2) / 2

    if media > 6 {
        return true
    } else if media < 6 {
        return false
    }

    // mata o programa
    panic("A média é exatamente 6")

}

func main() {
    fmt.Println(alunoEstaAprovado(6, 6))
}
