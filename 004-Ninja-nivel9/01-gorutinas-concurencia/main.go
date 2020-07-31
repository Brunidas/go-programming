package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("- - - - INICIO ")
	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera ")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hola desde la segunda")
		wg.Done()
	}()

	fmt.Println("- - - - MEDIO ")
	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas: %v\n", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("A punto de finalizar")

	fmt.Println("- - - - FINAL ")
	fmt.Printf("Numero de CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas: %v\n", runtime.NumGoroutine())
}
