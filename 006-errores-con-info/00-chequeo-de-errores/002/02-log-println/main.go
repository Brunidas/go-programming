package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("Un error ocurrio,", err)
	}

}
