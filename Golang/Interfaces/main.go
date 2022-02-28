package main

import (
    "fmt"
    "math"
)

type forma interface {
    area() float64
}

func escrever(f forma) {
    fmt.Printf("Aréa de %0.2f \n", f.area())
}

type retangulo struct {
    altura  float64
    largura float64
}

type circulo struct {
    raio float64
}

func (r retangulo) area() float64 {
    return r.altura * r.largura
}

func (c circulo) area() float64 {
    return math.Pi * (c.raio * c.raio)
}

func main() {
    r := retangulo{10, 15}
    escrever(r)

    c := circulo{80}
    escrever(c)
}
