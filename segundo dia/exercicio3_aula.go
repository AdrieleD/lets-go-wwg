package main

import "fmt"

func main() {

	/*Laços de repetição*/

	listaDeMercado := []string{"abacate", "sabonete", "azeite", "tomate", "banana"}

	for indice, valor := range listaDeMercado {
		fmt.Printf("%d - %s\n", indice+1, valor)
	}
}
