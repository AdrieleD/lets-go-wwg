package main

import (
	"fmt"
)

func main() {

	// # Exercício 01 - Extra
	var quilometros int
	quilometros = 150

	fmt.Println(quilometros)

	// Ocorre o erro "./prog.go:9:14: constant 150 overflows" pq o valor 150 ultrapassa o limite de int8 q é de -128 ~ 127,
	// para imprimir o valor é necessário trocar o tipo para "int"
}
