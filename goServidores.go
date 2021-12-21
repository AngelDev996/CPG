package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	canal := make(chan string)
	servidores := []string{
		"http://facebook.com",
		"http://google.com",
		"http://instagram.com",
		"http://fast.com",
	}

	for _, servidor := range servidores {
		go revisarServer(servidor, canal)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoPasado)
}

func revisarServer(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println("Servidor no disponible")
		canal <- servidor + " No se encuentra disponible"
	} else {
		fmt.Println("Servidor en funcionamiento")
		canal <- servidor + " Servidor disponible"

	}
}
