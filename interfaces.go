package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) mover() string {
	return "perro caminando"
}
func (pez) mover() string {
	return "pez nadando"

}
func (pajaro) mover() string {
	return "pajaro volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {

	pe := perro{}
	moverAnimal(pe)

	pa := pajaro{}
	moverAnimal(pa)

	pez := pez{}
	moverAnimal(pez)
}
