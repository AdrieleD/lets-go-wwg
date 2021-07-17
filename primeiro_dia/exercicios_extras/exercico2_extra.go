package main

import (
	"fmt"
	"time"
)

func main() {
	// # Exercício 02 - Extra
	dataAtual := time.Now().Year()
	var dataNascimento int
	fmt.Println("Digite o ano que vc nasceu:")
	fmt.Scanf("%d", &dataNascimento)
	idade := dataAtual - dataNascimento
	fmt.Printf("Você tem %d anos.", idade)
}
