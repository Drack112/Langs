package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 0:
		return "Coloque um número de 1 a 7"
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "oof"
	}
}

func diaDaSemana2(numero1 int) string {

	var diaDaSemana1 string

	switch {
	case numero1 == 0:
		diaDaSemana1 = "Coloque um número de 1 a 7"
		fallthrough
	case numero1 == 1:
		diaDaSemana1 = "Domingo"
	case numero1 == 2:
		diaDaSemana1 = "Segunda"
	case numero1 == 3:
		diaDaSemana1 = "Terça"
	case numero1 == 4:
		diaDaSemana1 = "Quarta"
	case numero1 == 5:
		diaDaSemana1 = "Quinta"
	case numero1 == 6:
		diaDaSemana1 = "Sexta"
	case numero1 == 7:
		diaDaSemana1 = "Sabado"
	default:
		diaDaSemana1 = "oof"
	}

	return diaDaSemana1
}

func main() {
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana2(0))
}
