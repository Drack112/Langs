package main

import "fmt"

func inverter(n1 int) int {
    return n1 * -1
}

func inverterPonteiro(n1 *int) {
    *n1 = *n1 * -1
}

func main() {
    numero := 20
    fmt.Println(inverter((numero)))

    n2 := 40
    inverterPonteiro(&n2)
    fmt.Println(n2)
}
