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

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nadar())
}
func moverPajaro(p pajaro) {
	fmt.Println(p.volar())
}

func main() {

	pe := perro{}
	moverPerro(pe)

	pa := pajaro{}
	moverPajaro(pa)

	pez := pez{}
	moverPez(pez)

}
