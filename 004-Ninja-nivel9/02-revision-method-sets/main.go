package main

import "fmt"

type persona struct {
	Nombre   string
	Apellido string
}

func (p *persona) hablar() {
	fmt.Println("Hola que tal, soy", p.Nombre, p.Apellido)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {

	p1 := persona{
		Nombre:   "Bruno",
		Apellido: "Fuentes",
	}
	//diAlgo(p1)
	// diAlgo(&p1)
	p1.hablar()
}
