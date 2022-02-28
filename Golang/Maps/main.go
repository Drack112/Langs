package main

import "fmt"

func main() {
	// Tipo dos valores
	usuario := map[string]string{
		// Tipo das chaves
		"nome":      "Pedro",
		"sobrenome": "Carvalho",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "ferreira",
		},
		"campus": {
			"nome":   "USP",
			"region": "SP",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Peixe",
	}
	fmt.Println(usuario2)
}
