package main

import (
	"fmt"
)

func main() {

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
}
