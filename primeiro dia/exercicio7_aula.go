package main

import (
	"fmt"
)

func main() {

	// # Exercício 07 - Aula
	ano := make(map[int]string)
	ano[1] = "Janeiro"
	ano[2] = "Fevereiro"
	ano[3] = "Março"
	ano[4] = "Abril"
	ano[5] = "Maio"
	ano[6] = "Junho"
	ano[7] = "Julho"
	ano[8] = "Agosto"
	ano[9] = "Setembro"
	ano[10] = "Outubro"
	ano[11] = "Novembro"
	ano[12] = "Dezembro"

	fmt.Printf("Mês 1 é %s e mês 12 é %s.\n", ano[1], ano[12])
	fmt.Println(len(ano))
}
