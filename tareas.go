package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) { //
	t.nombre = nombre
}
func main() {
	t := &task{ //t es un apuntador(apunta al espacio de memoria de task)
		nombre:      "Completar curso",
		descripcion: "Completacion de curso",
	}
	// fmt.Println(t)
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("Otro curso")
	t.actualizarDescripcion("Otra descripcion")
	fmt.Printf("%+v\n", t)

}
