package main

import "fmt"

func main() {

	/*Functions*/
	cumprimento := GereCumprimento("Maria", 19)
	fmt.Printf("%s\n", cumprimento)
}

func GereCumprimento(nome string, hora int) string {
	cumprimento := ""
	switch {
	case hora >= 6 && hora < 12:
		cumprimento = "Bom dia!"
	case hora >= 12 && hora < 18:
		cumprimento = "Boa Tarde!"
	case hora >= 18 && hora < 24:
		cumprimento = "Boa Noite!"
	case hora >= 0 && hora < 6:
		cumprimento = "Boa Madrugada!"
	default:
		cumprimento = "Tudo bem?"
	}

	frase := fmt.Sprintf("Olar, %s! %s\n", nome, cumprimento)
	return frase
}
