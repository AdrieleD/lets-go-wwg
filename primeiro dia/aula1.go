package main

import (
	"fmt"
)

func main() {

	// # Exercício 01
	a := 27
	b := 401
	c := 10
	d := 15
	fmt.Println(a, b, c, d)

	// # Exercício 02
	a := 230
	b := 27
	soma := a + b
	subtracao := a - b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(soma)
	fmt.Println(subtracao)

	// # Exercício 01 - Aula
	var pesoDoTomate, quantidadeDeGarrafasDeAzeite, unidadesDeSabonete float64
	pesoDoTomate = 2.600
	quantidadeDeGarrafasDeAzeite = 2
	unidadesDeSabonete = 7

	var precoDoTomate, precoDoAzeite, precoDoSabonete float64
	precoDoTomate = 10.3
	precoDoAzeite = 19
	precoDoSabonete = 5.80

	valorDaCompra := (pesoDoTomate * precoDoTomate) + (quantidadeDeGarrafasDeAzeite * precoDoAzeite) + (unidadesDeSabonete * precoDoSabonete)
	fmt.Printf("o valor da compra deu: \n%v \n\tsó uma garrafa de azeite já custou %v", valorDaCompra, precoDoAzeite)

	// # Exercício 02 - Aula
	nome := "Adriele"
	cor := "Preta"

	fmt.Printf("Olá, meu nome é %s e eu gosto da cor %s.", nome, cor)

	// # Exercício 03 - Aula
	a := 15 >= 15         // true
	b := 100 < 1000       // true
	c := 10 == 18         // true
	d := 5 != 5           // false
	e := 89 > 0 && 50 > 0 // true

	fmt.Printf("o tipo de a é %T e seu valor é %v\n", a, a)
	fmt.Printf("o tipo de b é %T e seu valor é %v\n", b, b)
	fmt.Printf("o tipo de c é %T e seu valor é %v\n", c, c)
	fmt.Printf("o tipo de d é %T e seu valor é %v\n", d, d)
	fmt.Printf("o tipo de e é %T e seu valor é %v\n", e, e)

	// # Exercício 04 - Aula
	// numeros := [6]string{"zero", "um", "dois", "tres", "quatro", "cinco"}
	fmt.Printf("o tipo é: %T\n", numeros)
	fmt.Println(numeros[:])

	// # Exercício 05 - Aula
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9, 10, 11, 12}
	slice3 := append(slice1, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// # Exercício 06 - Aula
	signos := make([]string, 12)

	signos[0] = "Áries"
	signos[1] = "Touro"
	signos[2] = "Gêmeos"
	signos[3] = "Câncer"
	signos[4] = "Leão"
	signos[5] = "Virgem"
	signos[6] = "Libra"
	signos[7] = "Escorpião"
	signos[8] = "Sagitário"
	signos[9] = "Capricórnio"
	signos[10] = "Aquário"
	signos[11] = "Peixes"

	fmt.Println(signos)
	fmt.Println(signos[1:4])

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
