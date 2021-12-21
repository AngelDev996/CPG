package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	servidores := []string{
		"http://facebook.com",
		"http://google.com",
		"http://instagram.com",
		"http://fast.com",
	}

	for _, seservidor := range servidores {
		revisarServer(seservidor)
	}
	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoPasado)
}

func revisarServer(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println("Servidor no disponible")
	} else {
		fmt.Println("Servidor en funcionamiento")
	}
}
