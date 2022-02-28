package main

import "fmt"

func a(posicao uint) uint {
    if posicao <= 1 {
        return posicao
    }
    return a(posicao-2) + a(posicao-1)
}

func main() {
    fmt.Println("Funções Recursivas")
    posicao := uint(12)

    for i := uint(0); i < posicao; i++ {
        fmt.Println(a(i))
    }

}
