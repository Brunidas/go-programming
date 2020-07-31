package main

import (
	"fmt"

	"gitlab.com/go-programming/010-Ninja-nivel13/02/cita"
	"gitlab.com/go-programming/010-Ninja-nivel13/02/palabra"
)

func main() {
	fmt.Println(palabra.Conteo(cita.Cuando))

	for k, v := range palabra.ConteoUso(cita.Cuando) {
		fmt.Println(v, k)
	}
}
