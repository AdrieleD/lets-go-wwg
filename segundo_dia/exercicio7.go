package main

import "fmt"

func main() {

	/*Laços de repetição*/

	hora := 0
	for hora < 24 {
		fmt.Printf("%.2d:00\n", hora)
		hora++
	}

}
