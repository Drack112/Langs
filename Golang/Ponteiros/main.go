package main

import "fmt"

func main() {
	n1 := 100
	n2 := n1

	fmt.Println(n1, n2)

	n1++
	fmt.Println(n1, n2)

	n1++
	n2 = n1
	fmt.Println(n1, n2)

	var n3 *int
	var ponteiro = &n3

	fmt.Println(n3, ponteiro)

	n4 := 100
	var ponteiro2 = &n4

	fmt.Println(n4, ponteiro2)
}
