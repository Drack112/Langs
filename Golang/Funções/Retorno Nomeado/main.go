package main

import "fmt"

func soma(x, y int8) (soma int8, sub int8) {
    soma = x + y
    sub = x - y
    return
}

func main() {
    fmt.Println(soma(10, 80))
}
