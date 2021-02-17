package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // sem pontinhos é um slice, não um array
	// dessa forma o compilador vai contar o array

	for i, numero := range numeros {
		/*
			vai percorrer o array recebendo cada item para cada indice
		*/

		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
