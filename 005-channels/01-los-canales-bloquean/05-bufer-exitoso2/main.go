package main

import "fmt"

func main() {
	// canal con buffer
	ca := make(chan int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)

}
