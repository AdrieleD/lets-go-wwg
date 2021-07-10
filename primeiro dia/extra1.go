package main

import (
	"fmt"
	"time"
)

func main() {

	// # Exercício 01 - Extra
	var quilometros int
	quilometros = 150

	fmt.Println(quilometros)

	// Ocorre o erro "./prog.go:9:14: constant 150 overflows" pq o valor 150 ultrapassa o limite de int8 q é de -128 ~ 127,
	// para imprimir o valor é necessário trocar o tipo para "int"

	// # Exercício 02 - Extra
	dataAtual := time.Now().Year()
	var dataNascimento int
	fmt.Println("Digite o ano que vc nasceu:")
	fmt.Scanf("%d", &dataNascimento)
	idade := dataAtual - dataNascimento
	fmt.Printf("Você tem %d anos.", idade)
}
