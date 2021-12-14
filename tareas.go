package main

import "fmt"

type tasklist struct {
	tareas []*task
}

func (t *tasklist) agregarALista(tl *task) {
	t.tareas = append(t.tareas, tl)
}
func (t *tasklist) eliminar(index int) { //metodo para eliminar un elemento de un slice, pasando como parametro el indice del slice
	t.tareas = append(t.tareas[:index], t.tareas[index+1:]...)
}

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
	t1 := &task{ //t es un apuntador(apunta al espacio de memoria de task)
		nombre:      "Completar curso",
		descripcion: "Completacion de curso",
	}

	t2 := &task{ //t es un apuntador(apunta al espacio de memoria de task)
		nombre:      "Completar curso2",
		descripcion: "Completacion de curso2",
	}

	listaTareas := &tasklist{
		tareas: []*task{
			t1, t2,
		},
	}
	fmt.Println(listaTareas.tareas[0])
	fmt.Println(len(listaTareas.tareas))
	listaTareas.eliminar(1)
	fmt.Println(len(listaTareas.tareas))
}
