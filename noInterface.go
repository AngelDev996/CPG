package main

import "fmt"

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) caminar() string {
	return "perro caminando"
}
func (pez) nadar() string {
	return "pez nadando"

}
func (pajaro) volar() string {
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
