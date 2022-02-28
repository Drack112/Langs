package main

import "fmt"

type usuario struct {
    name  string
    idade int
}

func (u usuario) salvar() {
    fmt.Printf("Usuário %s salvo no banco de dados.", u.name)
}

func main() {
    user := usuario{"Drack", 12}

    fmt.Println(user)
    user.salvar()
}
