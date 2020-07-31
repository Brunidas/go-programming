package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("El SO es:", runtime.GOOS)
	fmt.Println("La ARCH es:", runtime.GOARCH)

}
