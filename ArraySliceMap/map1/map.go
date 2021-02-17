package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Baiana"
	aprovados[98745214] = "Fellipe"
	aprovados[89522155] = "Rael"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 98745214)
	fmt.Println(aprovados[98745214])
}
