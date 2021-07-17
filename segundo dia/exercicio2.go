package main

import "fmt"

func main() {

	/*Condicionais*/

	// # Exemplo e exercício 3
	idade := 24
	if idade < 18 {
		fmt.Println("Menor de idade")
	} else if idade >= 18 && idade < 60 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Idoso")
	}

}
