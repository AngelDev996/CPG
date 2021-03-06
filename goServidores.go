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

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {
			go revisarServer(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoPasado)
}

func revisarServer(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + " No se encuentra disponible"
	} else {
		canal <- servidor + " Servidor disponible"

	}
}

