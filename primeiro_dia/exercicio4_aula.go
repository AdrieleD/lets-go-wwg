package main

import (
	"fmt"
)

func main() {
	// # Exercício 04 - Aula
	numeros := [6]string{"zero", "um", "dois", "tres", "quatro", "cinco"}
	fmt.Printf("o tipo é: %T\n", numeros)
	fmt.Println(numeros[:])
}
