package main

import "fmt"

func main() {
    /*
       func() {
           fmt.Println("Arroz")
       }()
    */
    /*
       func(texto string) {
           fmt.Println(texto)
       }("Batata")
    */
    f := func(texto string) string {
        return fmt.Sprintf("Recebi --> %s", texto)
    }("Texto")
    fmt.Println(f)
}
