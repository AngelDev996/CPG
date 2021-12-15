package main

import "fmt"

func main() {
	m1 := make(map[string]int) //Los maps son como diccionarios, se debe de declarar el tipo de la llave y el tipo del valor

	m1["a"] = 8
	fmt.Println(m1)
	fmt.Println(m1["a"])
}
