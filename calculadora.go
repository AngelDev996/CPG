package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operador1, err1 := strconv.Atoi(valores[0]) //Atoi castea string a enteros
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Hubo un error")

	} else {
		fmt.Println(operador1 + operador2)
	}

}
