package main

import "fmt"

func main() {
	// send-only
	ca := make(chan<- int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("-------")
	fmt.Printf("%T\n", ca)
}
