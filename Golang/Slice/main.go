package main

import "fmt"

func main() {
    array1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println(array1)

    array1 = append(array1, 9)
    fmt.Println(array1)

    slice := array1[0]
    fmt.Println(slice)

    slice2 := array1[0:5]
    fmt.Println(slice2)
}
