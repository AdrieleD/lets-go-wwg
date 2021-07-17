package main

import "fmt"

type Pessoa struct {
	nome         string
	idade        int
	corPreferida string
}

func main() {

	// # Exercício 01
	/* Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
	Crie três variáveis do tipo pessoa.
	Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.*/

	pessoa1 := Pessoa{
		nome:         "João",
		idade:        10,
		corPreferida: "Azul",
	}

	pessoa2 := Pessoa{
		nome:         "Maria",
		idade:        12,
		corPreferida: "Rosa",
	}

	pessoa3 := Pessoa{
		nome:         "Laura",
		idade:        8,
		corPreferida: "Vermelho",
	}

	fmt.Printf("O meu nome é %s, tenho %d anos e minha cor preferida é %s.\n", pessoa1.nome, pessoa1.idade, pessoa1.corPreferida)
	fmt.Printf("O meu nome é %s, tenho %d anos e minha cor preferida é %s.\n", pessoa2.nome, pessoa2.idade, pessoa2.corPreferida)
	fmt.Printf("O meu nome é %s, tenho %d anos e minha cor preferida é %s.\n", pessoa3.nome, pessoa3.idade, pessoa3.corPreferida)

}
