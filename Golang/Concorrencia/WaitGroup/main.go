package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // duas go rotines

	go func() {
		escrever("Hello World!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Lesgoo!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}
}
