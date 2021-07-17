package main

import "fmt"

type Apartamento struct {
	numero              int
	nomeDaProprietaria  string
	possuiVagaDeGaragem bool
}

func main() {

	ap1 := Apartamento{
		numero:              13010,
		nomeDaProprietaria:  "Talia",
		possuiVagaDeGaragem: true,
	}

	fmt.Printf("%+v\n", ap1)
	fmt.Printf("O apartamento %d tem como oproprietária a %s. Tem vaga de garagem? %t.\n", ap1.numero, ap1.nomeDaProprietaria, ap1.possuiVagaDeGaragem)

}
