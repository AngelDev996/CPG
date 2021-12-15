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

func (t *tasklist) imprimirLista() {
	for _, tarea := range t.tareas {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
		fmt.Println("Completado:", tarea.completado)
	}
}

func (t *tasklist) imprimirCompletos() {
	for _, tarea := range t.tareas {
		if tarea.completado == true {
			fmt.Println("Tareas completadas:", tarea.nombre)
		}
	}
}

func main() {
	t1 := &task{ //t es un apuntador(apunta al espacio de memoria de task)
		nombre:      "Completar curso1",
		descripcion: "Completacion de curso1",
	}

	t2 := &task{ //t es un apuntador(apunta al espacio de memoria de task)
		nombre:      "Completar curso2",
		descripcion: "Completacion de curso2",
	}

	t3 := &task{ //t es un apuntador(apunta al espacio de memoria de task)
		nombre:      "Completar curso3",
		descripcion: "Completacion de curso3",
		completado:  true,
	}

	listaTareas := &tasklist{
		tareas: []*task{
			t1, t2,
		},
	}
	listaTareas.agregarALista(t3)

	for i := 0; i < len(listaTareas.tareas); i++ { //ciclo for clasico
		fmt.Println("indice:", i, "titulo:", listaTareas.tareas[i].nombre)
	}

	for index, tarea := range listaTareas.tareas { // for mejorado
		fmt.Println(index, tarea.nombre)
	}

	for j := 0; j < 11; j++ { //uso de break
		if j == 5 {
			break
		}
		fmt.Println(j)

	}

	for j := 0; j < 11; j++ { //continue
		if j == 5 {
			continue
		}
		fmt.Println(j)

	}
	listaTareas.imprimirLista()
	listaTareas.imprimirCompletos()
}
