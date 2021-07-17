package main

import "fmt"

func main() {

	/*Functions*/
	soma := SomaNumeros(7, 5)
	fmt.Printf("Resultado da soma: %d\n", soma)
}

func SomaNumeros(a, b int) int {
	return a + b
}
