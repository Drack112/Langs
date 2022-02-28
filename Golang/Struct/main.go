package main

import "fmt"

type user struct {
	name  string
	idade int
	ativo bool
	langs
}

type langs struct {
	lang      string
	framework string
}

func main() {
	var u user

	u.name = "Drack"
	u.idade = 32
	u.ativo = true

	fmt.Println(u)

	usuario1 := user{name: "Rodrigo", idade: 52, ativo: false}
	fmt.Println(usuario1)

	usuario2 := user{ativo: false}
	fmt.Println(usuario2)

	usuario3 := user{name: "João", idade: 15, ativo: false, langs: langs{lang: "Python", framework: "Django"}}
	fmt.Println(usuario3)
}
