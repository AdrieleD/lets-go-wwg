package main

import "fmt"

func main() {

	/*Condicionais*/

	// # Execício 01

	anoDeNascimento := 2015
	anoAtual := 2021

	podeVotar := (anoAtual - anoDeNascimento) >= 16
	if podeVotar {
		fmt.Println("Esta pessoa pode votar.")
	} else {
		fmt.Println("Esta pessoa não pode votar.")
	}
}
