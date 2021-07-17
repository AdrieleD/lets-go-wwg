package main

import (
	"fmt"
)

func main() {

	// # Exercício 05 - Aula
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9, 10, 11, 12}
	slice3 := append(slice1, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
