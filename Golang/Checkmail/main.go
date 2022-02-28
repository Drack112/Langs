package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateFormat("oof@oof.com")

	if err != nil {
		fmt.Println("Email está no formato incorreto.")
	} else {
		fmt.Println("Email está no formato correto.")
	}
}
