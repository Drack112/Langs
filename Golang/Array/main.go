package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string
	array1[0] = "Melancia"
	array1[1] = "Kiwi"
	array1[2] = "Uva"
	array1[3] = "Banana"
	array1[4] = "Maçã"
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array4)

	fmt.Println("-------------------------")
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(array4))
	fmt.Println("-------------------------")

}
