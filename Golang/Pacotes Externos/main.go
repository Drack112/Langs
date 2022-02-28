package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	email := checkmail.ValidateFormat("admin@admin.com")
	fmt.Println(email) // nil pq n tem erro
}
