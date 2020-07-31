package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("El SO es :, %v\nArquitectura %v\n", runtime.GOOS, runtime.GOARCH )
}
